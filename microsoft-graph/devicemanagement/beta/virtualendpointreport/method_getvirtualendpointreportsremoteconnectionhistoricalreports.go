package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions() GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions {
	return GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsRemoteConnectionHistoricalReports - Invoke action getRemoteConnectionHistoricalReports. Get
// the remote connection history records of a Cloud PC during a given period. This report contains data such as
// signInDateTime, signOutDateTime, usageInHour, remoteSignInTimeInSec and roundTripTimeInMsP50, and so on. This data is
// aggregated hourly for a specified time period, such as the last seven days.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsRemoteConnectionHistoricalReports(ctx context.Context, input GetVirtualEndpointReportsRemoteConnectionHistoricalReportsRequest, options GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions) (result GetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getRemoteConnectionHistoricalReports",
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
