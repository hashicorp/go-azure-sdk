package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsConnectionQualityReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsConnectionQualityReportsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetVirtualEndpointReportsConnectionQualityReportsOperationOptions() GetVirtualEndpointReportsConnectionQualityReportsOperationOptions {
	return GetVirtualEndpointReportsConnectionQualityReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsConnectionQualityReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsConnectionQualityReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsConnectionQualityReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsConnectionQualityReports - Invoke action getConnectionQualityReports. Get the overall
// connection quality reports for all devices within a current tenant during a given time period, including metrics like
// the average round trip time (P50), average available bandwidth, and UDP connection percentage. Get also other
// real-time metrics such as last connection round trip time, last connection client IP, last connection gateway, and
// last connection protocol.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsConnectionQualityReports(ctx context.Context, input GetVirtualEndpointReportsConnectionQualityReportsRequest, options GetVirtualEndpointReportsConnectionQualityReportsOperationOptions) (result GetVirtualEndpointReportsConnectionQualityReportsOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/reports/getConnectionQualityReports",
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
