package permissionsanalyticaw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePermissionsAnalyticAwOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePermissionsAnalyticAwOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePermissionsAnalyticAwOperationOptions() DeletePermissionsAnalyticAwOperationOptions {
	return DeletePermissionsAnalyticAwOperationOptions{}
}

func (o DeletePermissionsAnalyticAwOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePermissionsAnalyticAwOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePermissionsAnalyticAwOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePermissionsAnalyticAw - Delete navigation property aws for identityGovernance
func (c PermissionsAnalyticAwClient) DeletePermissionsAnalyticAw(ctx context.Context, options DeletePermissionsAnalyticAwOperationOptions) (result DeletePermissionsAnalyticAwOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/aws",
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
