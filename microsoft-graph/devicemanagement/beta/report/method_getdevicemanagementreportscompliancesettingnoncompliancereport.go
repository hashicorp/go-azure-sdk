package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementReportsComplianceSettingNonComplianceReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetDeviceManagementReportsComplianceSettingNonComplianceReport ...
func (c ReportClient) GetDeviceManagementReportsComplianceSettingNonComplianceReport(ctx context.Context, input GetDeviceManagementReportsComplianceSettingNonComplianceReportRequest) (result GetDeviceManagementReportsComplianceSettingNonComplianceReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/reports/getComplianceSettingNonComplianceReport",
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
