package exportdevicemanagementvirtualendpointreportsjobs

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportVirtualEndpointReportsJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCExportJob
}

// ExportVirtualEndpointReportsJob - Create cloudPcExportJob. Create a new cloudPcExportJob resource to initiate
// downloading the entire or specified portion of a report. Use the GET cloudPcExportJob operation to verify the
// exportJobStatus property of the cloudPcExportJob resource. When the property result is completed, the report finishes
// downloading to the location specified by the exportUrl property.
func (c ExportDeviceManagementVirtualEndpointReportsJobsClient) ExportVirtualEndpointReportsJob(ctx context.Context, input beta.CloudPCExportJob) (result ExportVirtualEndpointReportsJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/virtualEndpoint/reports/exportJobs",
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

	var model beta.CloudPCExportJob
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
