package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsInaccessibleCloudPCReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// GetVirtualEndpointReportsInaccessibleCloudPCReport - Invoke action getInaccessibleCloudPcReports. Get inaccessible
// Cloud PCs with details, including the latest health state, failed connection count, failed health check count, and
// system status. An inaccessible Cloud PC represents a Cloud PC that is in an unavailable state (at least one of the
// health checks failed) or has consecutive user connections failure.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsInaccessibleCloudPCReport(ctx context.Context, input GetVirtualEndpointReportsInaccessibleCloudPCReportRequest) (result GetVirtualEndpointReportsInaccessibleCloudPCReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/getInaccessibleCloudPcReports",
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
