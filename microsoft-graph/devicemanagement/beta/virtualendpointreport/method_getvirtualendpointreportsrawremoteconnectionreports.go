package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsRawRemoteConnectionReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions() GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions {
	return GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsRawRemoteConnectionReports - Invoke action getRawRemoteConnectionReports. Get the raw
// real-time remote connection report for a Cloud PC without any calculation, such as roundTripTime or available
// bandwidth, which are aggregated hourly from the raw event data.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsRawRemoteConnectionReports(ctx context.Context, input GetVirtualEndpointReportsRawRemoteConnectionReportsRequest, options GetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions) (result GetVirtualEndpointReportsRawRemoteConnectionReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getRawRemoteConnectionReports",
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
