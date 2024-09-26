package crosstenantaccesspolicydefault

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions() ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions {
	return ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions{}
}

func (o ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetCrossTenantAccessPolicyDefaultToSystemDefault - Invoke action resetToSystemDefault. Reset any changes made to
// the default configuration in a cross-tenant access policy back to the system default.
func (c CrossTenantAccessPolicyDefaultClient) ResetCrossTenantAccessPolicyDefaultToSystemDefault(ctx context.Context, options ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationOptions) (result ResetCrossTenantAccessPolicyDefaultToSystemDefaultOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/default/resetToSystemDefault",
		RetryFunc:     options.RetryFunc,
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
