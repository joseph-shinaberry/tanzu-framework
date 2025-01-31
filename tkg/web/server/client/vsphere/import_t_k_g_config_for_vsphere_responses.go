// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ImportTKGConfigForVsphereReader is a Reader for the ImportTKGConfigForVsphere structure.
type ImportTKGConfigForVsphereReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportTKGConfigForVsphereReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportTKGConfigForVsphereOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportTKGConfigForVsphereBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportTKGConfigForVsphereUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportTKGConfigForVsphereInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportTKGConfigForVsphereOK creates a ImportTKGConfigForVsphereOK with default headers values
func NewImportTKGConfigForVsphereOK() *ImportTKGConfigForVsphereOK {
	return &ImportTKGConfigForVsphereOK{}
}

/*ImportTKGConfigForVsphereOK handles this case with default header values.

Generated TKG configuration successfully
*/
type ImportTKGConfigForVsphereOK struct {
	Payload *models.VsphereRegionalClusterParams
}

func (o *ImportTKGConfigForVsphereOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/config/import][%d] importTKGConfigForVsphereOK  %+v", 200, o.Payload)
}

func (o *ImportTKGConfigForVsphereOK) GetPayload() *models.VsphereRegionalClusterParams {
	return o.Payload
}

func (o *ImportTKGConfigForVsphereOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VsphereRegionalClusterParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForVsphereBadRequest creates a ImportTKGConfigForVsphereBadRequest with default headers values
func NewImportTKGConfigForVsphereBadRequest() *ImportTKGConfigForVsphereBadRequest {
	return &ImportTKGConfigForVsphereBadRequest{}
}

/*ImportTKGConfigForVsphereBadRequest handles this case with default header values.

Bad request
*/
type ImportTKGConfigForVsphereBadRequest struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForVsphereBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/config/import][%d] importTKGConfigForVsphereBadRequest  %+v", 400, o.Payload)
}

func (o *ImportTKGConfigForVsphereBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForVsphereBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForVsphereUnauthorized creates a ImportTKGConfigForVsphereUnauthorized with default headers values
func NewImportTKGConfigForVsphereUnauthorized() *ImportTKGConfigForVsphereUnauthorized {
	return &ImportTKGConfigForVsphereUnauthorized{}
}

/*ImportTKGConfigForVsphereUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ImportTKGConfigForVsphereUnauthorized struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForVsphereUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/config/import][%d] importTKGConfigForVsphereUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportTKGConfigForVsphereUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForVsphereUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForVsphereInternalServerError creates a ImportTKGConfigForVsphereInternalServerError with default headers values
func NewImportTKGConfigForVsphereInternalServerError() *ImportTKGConfigForVsphereInternalServerError {
	return &ImportTKGConfigForVsphereInternalServerError{}
}

/*ImportTKGConfigForVsphereInternalServerError handles this case with default header values.

Internal server error
*/
type ImportTKGConfigForVsphereInternalServerError struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForVsphereInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/config/import][%d] importTKGConfigForVsphereInternalServerError  %+v", 500, o.Payload)
}

func (o *ImportTKGConfigForVsphereInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForVsphereInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
