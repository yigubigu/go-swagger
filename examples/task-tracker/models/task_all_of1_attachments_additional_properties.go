// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskAllOf1AttachmentsAdditionalProperties task all of1 attachments additional properties
// swagger:model taskAllOf1AttachmentsAdditionalProperties
type TaskAllOf1AttachmentsAdditionalProperties struct {

	// The content type of the file.
	//
	// The content type of the file is inferred from the upload request.
	//
	// Read Only: true
	ContentType string `json:"contentType,omitempty"`

	// Extra information to attach to the file.
	//
	// This is a free form text field with support for github flavored markdown.
	//
	// Min Length: 3
	Description string `json:"description,omitempty"`

	// The name of the file.
	//
	// This name is inferred from the upload request.
	//
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The file size in bytes.
	//
	// This property was generated during the upload request of the file.
	// Read Only: true
	Size float64 `json:"size,omitempty"`

	// The url to download or view the file.
	//
	// This URL is generated on the server, based on where it was able to store the file when it was uploaded.
	//
	// Read Only: true
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this task all of1 attachments additional properties
func (m *TaskAllOf1AttachmentsAdditionalProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAllOf1AttachmentsAdditionalProperties) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 3); err != nil {
		return err
	}

	return nil
}

func (m *TaskAllOf1AttachmentsAdditionalProperties) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAllOf1AttachmentsAdditionalProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAllOf1AttachmentsAdditionalProperties) UnmarshalBinary(b []byte) error {
	var res TaskAllOf1AttachmentsAdditionalProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
