package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReport ...
func (c ReportClient) GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReport(ctx context.Context, input GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReportRequest) (result GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/reports/getDeviceConfigurationPolicySettingsSummaryReport",
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
