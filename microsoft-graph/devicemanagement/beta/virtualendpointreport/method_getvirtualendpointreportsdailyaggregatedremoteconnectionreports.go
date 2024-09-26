package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions() GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions {
	return GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReports - Invoke action
// getDailyAggregatedRemoteConnectionReports. Get the daily aggregated remote connection reports, such as round trip
// time, available bandwidth, and so on, in a given period.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReports(ctx context.Context, input GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsRequest, options GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions) (result GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getDailyAggregatedRemoteConnectionReports",
		RetryFunc:     options.RetryFunc,
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
