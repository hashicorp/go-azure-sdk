package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsActionStatusReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsActionStatusReportsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetVirtualEndpointReportsActionStatusReportsOperationOptions() GetVirtualEndpointReportsActionStatusReportsOperationOptions {
	return GetVirtualEndpointReportsActionStatusReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsActionStatusReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsActionStatusReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsActionStatusReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsActionStatusReports - Invoke action getActionStatusReports. Get the remote action status
// reports, including data such as the Cloud PC ID, Cloud PC device display name, initiating user's principal name,
// device owner's user principal name, action taken, and action state.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsActionStatusReports(ctx context.Context, input GetVirtualEndpointReportsActionStatusReportsRequest, options GetVirtualEndpointReportsActionStatusReportsOperationOptions) (result GetVirtualEndpointReportsActionStatusReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getActionStatusReports",
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
