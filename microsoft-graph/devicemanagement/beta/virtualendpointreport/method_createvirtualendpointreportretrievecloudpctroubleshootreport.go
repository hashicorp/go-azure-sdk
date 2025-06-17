package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions() CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions {
	return CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions{}
}

func (o CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReport - Invoke action retrieveCloudPcTroubleshootReports. Get
// troubleshooting reports for Cloud PCs. You can get a regional troubleshooting report, a report with troubleshooting
// details, a report with troubleshooting trends, or a report on the number of troubleshooting issues.
func (c VirtualEndpointReportClient) CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReport(ctx context.Context, input CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportRequest, options CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationOptions) (result CreateVirtualEndpointReportRetrieveCloudPCTroubleshootReportOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/reports/retrieveCloudPcTroubleshootReports",
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
