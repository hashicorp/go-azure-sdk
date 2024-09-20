package virtualendpointreportexportjob

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportExportJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCExportJob
}

type GetVirtualEndpointReportExportJobOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetVirtualEndpointReportExportJobOperationOptions() GetVirtualEndpointReportExportJobOperationOptions {
	return GetVirtualEndpointReportExportJobOperationOptions{}
}

func (o GetVirtualEndpointReportExportJobOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportExportJobOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetVirtualEndpointReportExportJobOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportExportJob - Get cloudPcExportJob. Read the properties and relationships of a cloudPcExportJob
// object. You can download a report by first creating a new cloudPcExportJob resource to initiate downloading. Use this
// GET operation to verify the exportJobStatus property of the cloudPcExportJob resource. The property becomes completed
// when the report finishes downloading in the location specified by the exportUrl property.
func (c VirtualEndpointReportExportJobClient) GetVirtualEndpointReportExportJob(ctx context.Context, id beta.DeviceManagementVirtualEndpointReportExportJobId, options GetVirtualEndpointReportExportJobOperationOptions) (result GetVirtualEndpointReportExportJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
