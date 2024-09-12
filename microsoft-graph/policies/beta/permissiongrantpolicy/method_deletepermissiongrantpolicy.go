package permissiongrantpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePermissionGrantPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePermissionGrantPolicyOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePermissionGrantPolicyOperationOptions() DeletePermissionGrantPolicyOperationOptions {
	return DeletePermissionGrantPolicyOperationOptions{}
}

func (o DeletePermissionGrantPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePermissionGrantPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePermissionGrantPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePermissionGrantPolicy - Delete permissionGrantPolicy. Delete a permissionGrantPolicy object.
func (c PermissionGrantPolicyClient) DeletePermissionGrantPolicy(ctx context.Context, id beta.PolicyPermissionGrantPolicyId, options DeletePermissionGrantPolicyOperationOptions) (result DeletePermissionGrantPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
