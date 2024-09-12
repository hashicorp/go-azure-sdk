package hardwareconfigurationrunsummary

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHardwareConfigurationRunSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwareConfigurationRunSummary
}

type GetHardwareConfigurationRunSummaryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetHardwareConfigurationRunSummaryOperationOptions() GetHardwareConfigurationRunSummaryOperationOptions {
	return GetHardwareConfigurationRunSummaryOperationOptions{}
}

func (o GetHardwareConfigurationRunSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHardwareConfigurationRunSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetHardwareConfigurationRunSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetHardwareConfigurationRunSummary - Get runSummary from deviceManagement. A summary of the results from an attempt
// to configure hardware settings. Read-Only.
func (c HardwareConfigurationRunSummaryClient) GetHardwareConfigurationRunSummary(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, options GetHardwareConfigurationRunSummaryOperationOptions) (result GetHardwareConfigurationRunSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/runSummary", id.ID()),
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

	var model beta.HardwareConfigurationRunSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
