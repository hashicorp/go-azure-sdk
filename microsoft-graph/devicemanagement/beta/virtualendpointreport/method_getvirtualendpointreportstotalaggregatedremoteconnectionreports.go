package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions() GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions {
	return GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReports - Invoke action
// getTotalAggregatedRemoteConnectionReports. Get the total aggregated remote connection usage of a Cloud PC during a
// given time span.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReports(ctx context.Context, input GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsRequest, options GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions) (result GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getTotalAggregatedRemoteConnectionReports",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
