package permissionsanalyticgcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePermissionsAnalyticGcpOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePermissionsAnalyticGcpOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeletePermissionsAnalyticGcpOperationOptions() DeletePermissionsAnalyticGcpOperationOptions {
	return DeletePermissionsAnalyticGcpOperationOptions{}
}

func (o DeletePermissionsAnalyticGcpOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePermissionsAnalyticGcpOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeletePermissionsAnalyticGcpOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePermissionsAnalyticGcp - Delete navigation property gcp for identityGovernance
func (c PermissionsAnalyticGcpClient) DeletePermissionsAnalyticGcp(ctx context.Context, options DeletePermissionsAnalyticGcpOperationOptions) (result DeletePermissionsAnalyticGcpOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/gcp",
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
