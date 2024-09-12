package assignmentfilter

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAssignmentFilterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceAndAppManagementAssignmentFilter
}

type GetAssignmentFilterOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAssignmentFilterOperationOptions() GetAssignmentFilterOperationOptions {
	return GetAssignmentFilterOperationOptions{}
}

func (o GetAssignmentFilterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAssignmentFilterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAssignmentFilterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAssignmentFilter - Get assignmentFilters from deviceManagement. The list of assignment filters
func (c AssignmentFilterClient) GetAssignmentFilter(ctx context.Context, id beta.DeviceManagementAssignmentFilterId, options GetAssignmentFilterOperationOptions) (result GetAssignmentFilterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDeviceAndAppManagementAssignmentFilterImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
