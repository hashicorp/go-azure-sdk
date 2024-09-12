package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport - Invoke action
// retrieveCrossRegionDisasterRecoveryReport. Retrieve the Windows 365 cross-region disaster recovery report, including
// CloudPcId, UserId, DeviceId, CloudPCDeviceDisplayName, UserPrincipalName, IsCrossRegionEnabled,
// CrossRegionHealthStatus, LicenseType, DisasterRecoveryStatus, CurrentRestorePointDateTime, and
// ActivationExpirationDateTime.
func (c VirtualEndpointReportClient) CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport(ctx context.Context, input CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportRequest) (result CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/retrieveCrossRegionDisasterRecoveryReport",
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
