package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// WorkflowGraph workflow graph
// swagger:model workflow_graph
type WorkflowGraph struct {

	// friendly name
	FriendlyName string `json:"friendlyName,omitempty"`

	// injectable name
	InjectableName string `json:"injectableName,omitempty"`

	// options
	Options interface{} `json:"options,omitempty"`

	// tasks
	Tasks []*WorkflowGraphTask `json:"tasks,omitempty"`
}

// Validate validates this workflow graph
func (m *WorkflowGraph) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowGraph) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {

		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {

			if err := m.Tasks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
