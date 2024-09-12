package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsConnectionQualityReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// GetVirtualEndpointReportsConnectionQualityReport - Invoke action getConnectionQualityReports. Get the overall
// connection quality reports for all devices within a current tenant during a given time period, including metrics like
// the average round trip time (P50), average available bandwidth, and UDP connection percentage. Get also other
// real-time metrics such as last connection round trip time, last connection client IP, last connection gateway, and
// last connection protocol.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsConnectionQualityReport(ctx context.Context, input GetVirtualEndpointReportsConnectionQualityReportRequest) (result GetVirtualEndpointReportsConnectionQualityReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/getConnectionQualityReports",
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
