// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DataLossPreventionPolicyRequestBuilder is request builder for DataLossPreventionPolicy
type DataLossPreventionPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataLossPreventionPolicyRequest
func (b *DataLossPreventionPolicyRequestBuilder) Request() *DataLossPreventionPolicyRequest {
	return &DataLossPreventionPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataLossPreventionPolicyRequest is request for DataLossPreventionPolicy
type DataLossPreventionPolicyRequest struct{ BaseRequest }

// Get performs GET request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Get(ctx context.Context) (resObj *DataLossPreventionPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Update(ctx context.Context, reqObj *DataLossPreventionPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}