package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReport - Invoke action
// getDailyAggregatedRemoteConnectionReports. Get the daily aggregated remote connection reports, such as round trip
// time, available bandwidth, and so on, in a given period.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReport(ctx context.Context, input GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportRequest) (result GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/getDailyAggregatedRemoteConnectionReports",
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
