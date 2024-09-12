package virtualendpointreport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVirtualEndpointReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteVirtualEndpointReportOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteVirtualEndpointReportOperationOptions() DeleteVirtualEndpointReportOperationOptions {
	return DeleteVirtualEndpointReportOperationOptions{}
}

func (o DeleteVirtualEndpointReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteVirtualEndpointReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteVirtualEndpointReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteVirtualEndpointReport - Delete navigation property reports for deviceManagement
func (c VirtualEndpointReportClient) DeleteVirtualEndpointReport(ctx context.Context, options DeleteVirtualEndpointReportOperationOptions) (result DeleteVirtualEndpointReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports",
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

	return
}
