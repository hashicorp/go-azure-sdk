package permissiongrantpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionGrantPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionGrantPolicy
}

type CreatePermissionGrantPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePermissionGrantPolicyOperationOptions() CreatePermissionGrantPolicyOperationOptions {
	return CreatePermissionGrantPolicyOperationOptions{}
}

func (o CreatePermissionGrantPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionGrantPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionGrantPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionGrantPolicy - Create permissionGrantPolicy. Creates a permissionGrantPolicy. A permission grant
// policy is used to describe the conditions under which permissions can be granted (for example, during application
// consent). After creating the permission grant policy, you can add include condition sets to add matching rules, and
// add exclude condition sets to add exclusion rules.
func (c PermissionGrantPolicyClient) CreatePermissionGrantPolicy(ctx context.Context, input beta.PermissionGrantPolicy, options CreatePermissionGrantPolicyOperationOptions) (result CreatePermissionGrantPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/policies/permissionGrantPolicies",
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

	var model beta.PermissionGrantPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
