package deviceconfigurationconflictsummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceConfigurationConflictSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceConfigurationConflictSummary
}

type GetDeviceConfigurationConflictSummaryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceConfigurationConflictSummaryOperationOptions() GetDeviceConfigurationConflictSummaryOperationOptions {
	return GetDeviceConfigurationConflictSummaryOperationOptions{}
}

func (o GetDeviceConfigurationConflictSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceConfigurationConflictSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceConfigurationConflictSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceConfigurationConflictSummary - Get deviceConfigurationConflictSummary from deviceManagement. Summary of
// policies in conflict state for this account.
func (c DeviceConfigurationConflictSummaryClient) GetDeviceConfigurationConflictSummary(ctx context.Context, id beta.DeviceManagementDeviceConfigurationConflictSummaryId, options GetDeviceConfigurationConflictSummaryOperationOptions) (result GetDeviceConfigurationConflictSummaryOperationResponse, err error) {
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

	var model beta.DeviceConfigurationConflictSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
