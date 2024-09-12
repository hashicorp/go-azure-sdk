package permissiongrantpolicyexclude

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPermissionGrantPolicyExcludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionGrantConditionSet
}

type GetPermissionGrantPolicyExcludeOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPermissionGrantPolicyExcludeOperationOptions() GetPermissionGrantPolicyExcludeOperationOptions {
	return GetPermissionGrantPolicyExcludeOperationOptions{}
}

func (o GetPermissionGrantPolicyExcludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPermissionGrantPolicyExcludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPermissionGrantPolicyExcludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPermissionGrantPolicyExclude - Get excludes from policies. Condition sets that are excluded in this permission
// grant policy. Automatically expanded on GET.
func (c PermissionGrantPolicyExcludeClient) GetPermissionGrantPolicyExclude(ctx context.Context, id beta.PolicyPermissionGrantPolicyIdExcludeId, options GetPermissionGrantPolicyExcludeOperationOptions) (result GetPermissionGrantPolicyExcludeOperationResponse, err error) {
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

	var model beta.PermissionGrantConditionSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
