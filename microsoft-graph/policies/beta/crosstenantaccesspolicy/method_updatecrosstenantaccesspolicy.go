package crosstenantaccesspolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCrossTenantAccessPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCrossTenantAccessPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCrossTenantAccessPolicyOperationOptions() UpdateCrossTenantAccessPolicyOperationOptions {
	return UpdateCrossTenantAccessPolicyOperationOptions{}
}

func (o UpdateCrossTenantAccessPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCrossTenantAccessPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCrossTenantAccessPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCrossTenantAccessPolicy - Update crossTenantAccessPolicy. Update the properties of a cross-tenant access
// policy.
func (c CrossTenantAccessPolicyClient) UpdateCrossTenantAccessPolicy(ctx context.Context, input beta.CrossTenantAccessPolicy, options UpdateCrossTenantAccessPolicyOperationOptions) (result UpdateCrossTenantAccessPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy",
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

	return
}
