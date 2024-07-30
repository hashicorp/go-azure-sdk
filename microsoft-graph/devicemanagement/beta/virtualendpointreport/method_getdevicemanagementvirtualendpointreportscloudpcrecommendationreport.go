package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReport ...
func (c VirtualEndpointReportClient) GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReport(ctx context.Context, input GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReportRequest) (result GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/getCloudPcRecommendationReports",
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

	return
}
