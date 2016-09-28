package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/consul"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
	"github.com/hashicorp/go-cleanhttp"
)

// Responses is an array of Response structs
type Responses []Response

// Server is the proxy server struct
type Server struct {
	Address string
	Port    int
	Store   *watcher.Monitor
	wg      *sync.WaitGroup
}

// Serve starts the Server on address:port and handles the routes
func (p *Server) Serve() {
	m := http.NewServeMux()
	m.HandleFunc("/test", p.HandleTest)
	m.HandleFunc("/", p.HandleNodes)
	http.ListenAndServe(fmt.Sprintf("%s:%d", p.Address, p.Port), m)
}

// NewServer initializes a new Server
func NewServer(proxyIP, serviceName, datacenter, backendAddr string, backend regStore.Backend, proxyPort int) (*Server, error) {
	consul.Register()
	m, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, backend)
	if err != nil {
		return nil, err
	}
	proxyServer := &Server{
		Address: proxyIP,
		Port:    proxyPort,
		Store:   m,
		wg:      &sync.WaitGroup{},
	}
	return proxyServer, nil
}

// HandleTest is....well a test
func (p *Server) HandleTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HandleNodes sends, recieves, and processes all the data
func (p *Server) HandleNodes(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//	fmt.Printf("Request Recieved => %s\n", time.Now())
	addrMap, err := p.GetAddresses(w, r)
	if len(addrMap) == 0 {
		w.WriteHeader(200)
		w.Write([]byte("[]"))
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	if (r.Method != "GET") && (len(addrMap) > 1) {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintln("Unsupported api call to multiple hosts. Use query string method."))) //TODO make this json
		return
	}
	ar := p.GetResp(r, addrMap)
	p.RespCheck(r, w, ar)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

// GetResp makes channels for the response and errors from http.Get.
// A go func is spun up for each http.Get and the responses are fed
// into their respective channels.
func (p *Server) GetResp(r *http.Request, addrs map[string]struct{}) Responses {
	cr := make(chan *Response, len(addrs))
	for entry := range addrs {
		p.wg.Add(1)
		go func(entry string, r *http.Request) {
			defer p.wg.Done()
			req, err := NewRequest(r, entry)
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			client := cleanhttp.DefaultClient()
			respGet, err := client.Do(req)
			//			fmt.Printf("Request Sent => %s\n", time.Now())
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			defer respGet.Body.Close()
			responseCopy, err := NewResponse(respGet)
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			cr <- responseCopy
		}(entry, r)
	}
	p.wg.Wait()
	close(cr)
	var ar Responses
	for entry := range cr {
		ar = append(ar, *entry)
	}
	return ar
}

// GetAddresses decides from where to retrieve the addresses
func (p *Server) GetAddresses(w http.ResponseWriter, r *http.Request) (map[string]struct{}, error) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	querySlice := r.URL.Query()
	if len(querySlice["ip"]) > 0 {
		addrMap := p.GetQueryAddresses(querySlice["ip"])
		return addrMap, nil
	}
	addrMap, err := p.GetStoredAddresses()
	if err != nil {
		return nil, err
	}
	return addrMap, nil
}

// GetStoredAddresses calls GetAddresses and returns a map of addresses
func (p *Server) GetStoredAddresses() (map[string]struct{}, error) {
	addresses, err := p.Store.GetAddresses()
	if err != nil {
		log.Printf("Did not get IP List ==> %s\n", err)
	}
	return addresses, err
}

// GetQueryAddresses retrives a url flag and returns a map of address(es)
func (p *Server) GetQueryAddresses(querySlice []string) map[string]struct{} {
	fmt.Println(querySlice)
	queryMap := make(map[string]struct{})
	for _, elem := range querySlice {
		ip, port, err := net.SplitHostPort(elem)
		if err != nil {
			log.Printf("Invalid port => %s\n", err)
			return nil
		}
		if net.ParseIP(ip) != nil {
			queryMap[ip+":"+port] = struct{}{}
		}
	}
	return queryMap
}

// RespCheck identifies the type of initialResp.Body and calls the appropriate
// helper function to write to the ResponseWriter.
func (p *Server) RespCheck(r *http.Request, w http.ResponseWriter, ar Responses) {
	w.Write([]byte("["))
	var cutSize, status int
	status = 500
	for i, r := range ar {
		if r.Body == nil {
			continue
		}
		if r.Body[0] == '[' {
			cutSize = 1
		} else if r.Body[0] == '{' {
			cutSize = 0
		} else {
			continue
		}
		w.Write(r.Body[cutSize : len(r.Body)-cutSize])
		if i != len(ar)-1 {
			w.Write([]byte(","))
		}
		if r.StatusCode < status {
			status = r.StatusCode
		}
	}
	if len(ar) <= 1 {
		for k, v := range ar[0].Header {
			for _, value := range v {
				w.Header().Set(k, value)
			}
		}
	}
	w.Write([]byte("]"))
	w.WriteHeader(status)
	//	fmt.Printf("Response Written => %s\n", time.Now())
}
