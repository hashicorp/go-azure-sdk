package crosstenantaccesspolicydefault

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCrossTenantAccessPolicyDefaultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CrossTenantAccessPolicyConfigurationDefault
}

type GetCrossTenantAccessPolicyDefaultOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetCrossTenantAccessPolicyDefaultOperationOptions() GetCrossTenantAccessPolicyDefaultOperationOptions {
	return GetCrossTenantAccessPolicyDefaultOperationOptions{}
}

func (o GetCrossTenantAccessPolicyDefaultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCrossTenantAccessPolicyDefaultOperationOptions) ToOData() *odata.Query {
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

func (o GetCrossTenantAccessPolicyDefaultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCrossTenantAccessPolicyDefault - Get crossTenantAccessPolicyConfigurationDefault. Read the default configuration
// of a cross-tenant access policy. This default configuration may be the service default assigned by Microsoft Entra ID
// (isServiceDefault is true) or may be customized in your tenant (isServiceDefault is false).
func (c CrossTenantAccessPolicyDefaultClient) GetCrossTenantAccessPolicyDefault(ctx context.Context, options GetCrossTenantAccessPolicyDefaultOperationOptions) (result GetCrossTenantAccessPolicyDefaultOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/default",
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

	var model stable.CrossTenantAccessPolicyConfigurationDefault
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
