// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MobileAppCollectionHasPayloadLinksRequestParameter undocumented
type MobileAppCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// MobileAppAssignRequestParameter undocumented
type MobileAppAssignRequestParameter struct {
	// MobileAppAssignments undocumented
	MobileAppAssignments []MobileAppAssignment `json:"mobileAppAssignments,omitempty"`
}

// MobileAppUpdateRelationshipsRequestParameter undocumented
type MobileAppUpdateRelationshipsRequestParameter struct {
	// Relationships undocumented
	Relationships []MobileAppRelationship `json:"relationships,omitempty"`
}

//
type MobileAppCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) HasPayloadLinks(reqObj *MobileAppCollectionHasPayloadLinksRequestParameter) *MobileAppCollectionHasPayloadLinksRequestBuilder {
	bb := &MobileAppCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *MobileAppCollectionHasPayloadLinksRequestBuilder) Request() *MobileAppCollectionHasPayloadLinksRequest {
	return &MobileAppCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *MobileAppCollectionHasPayloadLinksRequest) Get(ctx context.Context) ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type MobileAppAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *MobileAppRequestBuilder) Assign(reqObj *MobileAppAssignRequestParameter) *MobileAppAssignRequestBuilder {
	bb := &MobileAppAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppAssignRequest struct{ BaseRequest }

//
func (b *MobileAppAssignRequestBuilder) Request() *MobileAppAssignRequest {
	return &MobileAppAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type MobileAppUpdateRelationshipsRequestBuilder struct{ BaseRequestBuilder }

// UpdateRelationships action undocumented
func (b *MobileAppRequestBuilder) UpdateRelationships(reqObj *MobileAppUpdateRelationshipsRequestParameter) *MobileAppUpdateRelationshipsRequestBuilder {
	bb := &MobileAppUpdateRelationshipsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateRelationships"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppUpdateRelationshipsRequest struct{ BaseRequest }

//
func (b *MobileAppUpdateRelationshipsRequestBuilder) Request() *MobileAppUpdateRelationshipsRequest {
	return &MobileAppUpdateRelationshipsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppUpdateRelationshipsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}