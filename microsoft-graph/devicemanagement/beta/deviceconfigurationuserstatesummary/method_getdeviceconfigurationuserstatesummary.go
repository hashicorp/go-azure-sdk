package deviceconfigurationuserstatesummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceConfigurationUserStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceConfigurationUserStateSummary
}

type GetDeviceConfigurationUserStateSummaryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceConfigurationUserStateSummaryOperationOptions() GetDeviceConfigurationUserStateSummaryOperationOptions {
	return GetDeviceConfigurationUserStateSummaryOperationOptions{}
}

func (o GetDeviceConfigurationUserStateSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceConfigurationUserStateSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceConfigurationUserStateSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceConfigurationUserStateSummary - Get deviceConfigurationUserStateSummaries from deviceManagement. The device
// configuration user state summary for this account.
func (c DeviceConfigurationUserStateSummaryClient) GetDeviceConfigurationUserStateSummary(ctx context.Context, options GetDeviceConfigurationUserStateSummaryOperationOptions) (result GetDeviceConfigurationUserStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceConfigurationUserStateSummaries",
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

	var model beta.DeviceConfigurationUserStateSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
