// Package admin provides access to the Admin Data Transfer API.
//
// See https://developers.google.com/admin-sdk/data-transfer/
//
// Usage example:
//
//   import "google.golang.org/api/admin/datatransfer/v1"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "admin:datatransfer_v1"
const apiName = "admin"
const apiVersion = "datatransfer_v1"
const basePath = "https://www.googleapis.com/admin/datatransfer/v1/"

// OAuth2 scopes used by this API.
const (
	// View and manage data transfers between users in your organization
	AdminDatatransferScope = "https://www.googleapis.com/auth/admin.datatransfer"

	// View data transfers between users in your organization
	AdminDatatransferReadonlyScope = "https://www.googleapis.com/auth/admin.datatransfer.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Applications = NewApplicationsService(s)
	s.Transfers = NewTransfersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Applications *ApplicationsService

	Transfers *TransfersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewApplicationsService(s *Service) *ApplicationsService {
	rs := &ApplicationsService{s: s}
	return rs
}

type ApplicationsService struct {
	s *Service
}

func NewTransfersService(s *Service) *TransfersService {
	rs := &TransfersService{s: s}
	return rs
}

type TransfersService struct {
	s *Service
}

// Application: The JSON template for an Application resource.
type Application struct {
	// Etag: Etag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The application's ID.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies the resource as a DataTransfer Application Resource.
	Kind string `json:"kind,omitempty"`

	// Name: The application's name.
	Name string `json:"name,omitempty"`

	// TransferParams: The list of all possible transfer parameters for this
	// application. These parameters can be used to select the data of the
	// user in this application to be transfered.
	TransferParams []*ApplicationTransferParam `json:"transferParams,omitempty"`
}

// ApplicationDataTransfer: Template to map fields of
// ApplicationDataTransfer resource.
type ApplicationDataTransfer struct {
	// ApplicationId: The application's ID.
	ApplicationId int64 `json:"applicationId,omitempty,string"`

	// ApplicationTransferParams: The transfer parameters for the
	// application. These parameters are used to select the data which will
	// get transfered in context of this application.
	ApplicationTransferParams []*ApplicationTransferParam `json:"applicationTransferParams,omitempty"`

	// ApplicationTransferStatus: Current status of transfer for this
	// application. (Read-only)
	ApplicationTransferStatus string `json:"applicationTransferStatus,omitempty"`
}

// ApplicationTransferParam: Template for application transfer
// parameters.
type ApplicationTransferParam struct {
	// Key: The type of the transfer parameter. eg: 'PRIVACY_LEVEL'
	Key string `json:"key,omitempty"`

	// Value: The value of the coressponding transfer parameter. eg:
	// 'PRIVATE' or 'SHARED'
	Value []string `json:"value,omitempty"`
}

// ApplicationsListResponse: Template for a collection of Applications.
type ApplicationsListResponse struct {
	// Applications: List of applications that support data transfer and are
	// also installed for the customer.
	Applications []*Application `json:"applications,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Identifies the resource as a collection of Applications.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token which will be used to specify next
	// page in list API.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// DataTransfer: The JSON template for a DataTransfer resource.
type DataTransfer struct {
	// ApplicationDataTransfers: List of per application data transfer
	// resources. It contains data transfer details of the applications
	// associated with this transfer resource. Note that this list is also
	// used to specify the applications for which data transfer has to be
	// done at the time of the transfer resource creation.
	ApplicationDataTransfers []*ApplicationDataTransfer `json:"applicationDataTransfers,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The transfer's ID (Read-only).
	Id string `json:"id,omitempty"`

	// Kind: Identifies the resource as a DataTransfer request.
	Kind string `json:"kind,omitempty"`

	// NewOwnerUserId: ID of the user to whom the data is being transfered.
	NewOwnerUserId string `json:"newOwnerUserId,omitempty"`

	// OldOwnerUserId: ID of the user whose data is being transfered.
	OldOwnerUserId string `json:"oldOwnerUserId,omitempty"`

	// OverallTransferStatusCode: Overall transfer status (Read-only).
	OverallTransferStatusCode string `json:"overallTransferStatusCode,omitempty"`

	// RequestTime: The time at which the data transfer was requested
	// (Read-only).
	RequestTime string `json:"requestTime,omitempty"`
}

// DataTransfersListResponse: Template for a collection of DataTransfer
// resources.
type DataTransfersListResponse struct {
	// DataTransfers: List of data transfer requests.
	DataTransfers []*DataTransfer `json:"dataTransfers,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Identifies the resource as a collection of data transfer
	// requests.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token which will be used to specify next
	// page in list API.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "datatransfer.applications.get":

type ApplicationsGetCall struct {
	s             *Service
	applicationId int64
	opt_          map[string]interface{}
}

// Get: Retrieves information about an application for the given
// application ID.
func (r *ApplicationsService) Get(applicationId int64) *ApplicationsGetCall {
	c := &ApplicationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ApplicationsGetCall) Fields(s ...googleapi.Field) *ApplicationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ApplicationsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications/{applicationId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"applicationId": strconv.FormatInt(c.applicationId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ApplicationsGetCall) Do() (*Application, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Application
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about an application for the given application ID.",
	//   "httpMethod": "GET",
	//   "id": "datatransfer.applications.get",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "ID of the application resource to be retrieved.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}",
	//   "response": {
	//     "$ref": "Application"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.datatransfer",
	//     "https://www.googleapis.com/auth/admin.datatransfer.readonly"
	//   ]
	// }

}

// method id "datatransfer.applications.list":

type ApplicationsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists the applications available for data transfer for a
// customer.
func (r *ApplicationsService) List() *ApplicationsListCall {
	c := &ApplicationsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// CustomerId sets the optional parameter "customerId": Immutable ID of
// the Google Apps account.
func (c *ApplicationsListCall) CustomerId(customerId string) *ApplicationsListCall {
	c.opt_["customerId"] = customerId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100.
func (c *ApplicationsListCall) MaxResults(maxResults int64) *ApplicationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list.
func (c *ApplicationsListCall) PageToken(pageToken string) *ApplicationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ApplicationsListCall) Fields(s ...googleapi.Field) *ApplicationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ApplicationsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["customerId"]; ok {
		params.Set("customerId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ApplicationsListCall) Do() (*ApplicationsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ApplicationsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the applications available for data transfer for a customer.",
	//   "httpMethod": "GET",
	//   "id": "datatransfer.applications.list",
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable ID of the Google Apps account.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications",
	//   "response": {
	//     "$ref": "ApplicationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.datatransfer",
	//     "https://www.googleapis.com/auth/admin.datatransfer.readonly"
	//   ]
	// }

}

// method id "datatransfer.transfers.get":

type TransfersGetCall struct {
	s              *Service
	dataTransferId string
	opt_           map[string]interface{}
}

// Get: Retrieves a data transfer request by its resource ID.
func (r *TransfersService) Get(dataTransferId string) *TransfersGetCall {
	c := &TransfersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.dataTransferId = dataTransferId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TransfersGetCall) Fields(s ...googleapi.Field) *TransfersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TransfersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "transfers/{dataTransferId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"dataTransferId": c.dataTransferId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TransfersGetCall) Do() (*DataTransfer, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DataTransfer
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a data transfer request by its resource ID.",
	//   "httpMethod": "GET",
	//   "id": "datatransfer.transfers.get",
	//   "parameterOrder": [
	//     "dataTransferId"
	//   ],
	//   "parameters": {
	//     "dataTransferId": {
	//       "description": "ID of the resource to be retrieved. This is returned in the response from the insert method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "transfers/{dataTransferId}",
	//   "response": {
	//     "$ref": "DataTransfer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.datatransfer",
	//     "https://www.googleapis.com/auth/admin.datatransfer.readonly"
	//   ]
	// }

}

// method id "datatransfer.transfers.insert":

type TransfersInsertCall struct {
	s            *Service
	datatransfer *DataTransfer
	opt_         map[string]interface{}
}

// Insert: Inserts a data transfer request.
func (r *TransfersService) Insert(datatransfer *DataTransfer) *TransfersInsertCall {
	c := &TransfersInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.datatransfer = datatransfer
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TransfersInsertCall) Fields(s ...googleapi.Field) *TransfersInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TransfersInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.datatransfer)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "transfers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TransfersInsertCall) Do() (*DataTransfer, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DataTransfer
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a data transfer request.",
	//   "httpMethod": "POST",
	//   "id": "datatransfer.transfers.insert",
	//   "path": "transfers",
	//   "request": {
	//     "$ref": "DataTransfer"
	//   },
	//   "response": {
	//     "$ref": "DataTransfer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.datatransfer"
	//   ]
	// }

}

// method id "datatransfer.transfers.list":

type TransfersListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists the transfers for a customer by source user, destination
// user, or status.
func (r *TransfersService) List() *TransfersListCall {
	c := &TransfersListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// CustomerId sets the optional parameter "customerId": Immutable ID of
// the Google Apps account.
func (c *TransfersListCall) CustomerId(customerId string) *TransfersListCall {
	c.opt_["customerId"] = customerId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100.
func (c *TransfersListCall) MaxResults(maxResults int64) *TransfersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// NewOwnerUserId sets the optional parameter "newOwnerUserId":
// Destination user's profile ID.
func (c *TransfersListCall) NewOwnerUserId(newOwnerUserId string) *TransfersListCall {
	c.opt_["newOwnerUserId"] = newOwnerUserId
	return c
}

// OldOwnerUserId sets the optional parameter "oldOwnerUserId": Source
// user's profile ID.
func (c *TransfersListCall) OldOwnerUserId(oldOwnerUserId string) *TransfersListCall {
	c.opt_["oldOwnerUserId"] = oldOwnerUserId
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// the next page in the list.
func (c *TransfersListCall) PageToken(pageToken string) *TransfersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Status sets the optional parameter "status": Status of the transfer.
func (c *TransfersListCall) Status(status string) *TransfersListCall {
	c.opt_["status"] = status
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TransfersListCall) Fields(s ...googleapi.Field) *TransfersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TransfersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["customerId"]; ok {
		params.Set("customerId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["newOwnerUserId"]; ok {
		params.Set("newOwnerUserId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["oldOwnerUserId"]; ok {
		params.Set("oldOwnerUserId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "transfers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TransfersListCall) Do() (*DataTransfersListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DataTransfersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the transfers for a customer by source user, destination user, or status.",
	//   "httpMethod": "GET",
	//   "id": "datatransfer.transfers.list",
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable ID of the Google Apps account.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "newOwnerUserId": {
	//       "description": "Destination user's profile ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "oldOwnerUserId": {
	//       "description": "Source user's profile ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify the next page in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "Status of the transfer.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "transfers",
	//   "response": {
	//     "$ref": "DataTransfersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.datatransfer",
	//     "https://www.googleapis.com/auth/admin.datatransfer.readonly"
	//   ]
	// }

}
