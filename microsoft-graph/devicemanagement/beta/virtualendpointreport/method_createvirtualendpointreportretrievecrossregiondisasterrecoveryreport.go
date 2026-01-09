package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions() CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions {
	return CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions{}
}

func (o CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport - Invoke action
// retrieveCrossRegionDisasterRecoveryReport. Retrieve the Windows 365 cross-region disaster recovery report, including
// cloudPcId, userId, deviceId, cloudPCDeviceDisplayName, userPrincipalName, enabledDRType, disasterRecoveryStatus,
// licenseType, drHealthStatus, currentRestorePointDateTime, backupCloudPcStatus, and activationExpirationDateTime.
func (c VirtualEndpointReportClient) CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport(ctx context.Context, input CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportRequest, options CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions) (result CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/reports/retrieveCrossRegionDisasterRecoveryReport",
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
