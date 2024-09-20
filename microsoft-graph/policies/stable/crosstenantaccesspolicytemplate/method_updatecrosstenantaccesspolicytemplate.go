package crosstenantaccesspolicytemplate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCrossTenantAccessPolicyTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCrossTenantAccessPolicyTemplateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCrossTenantAccessPolicyTemplateOperationOptions() UpdateCrossTenantAccessPolicyTemplateOperationOptions {
	return UpdateCrossTenantAccessPolicyTemplateOperationOptions{}
}

func (o UpdateCrossTenantAccessPolicyTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCrossTenantAccessPolicyTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCrossTenantAccessPolicyTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCrossTenantAccessPolicyTemplate - Update the navigation property templates in policies
func (c CrossTenantAccessPolicyTemplateClient) UpdateCrossTenantAccessPolicyTemplate(ctx context.Context, input stable.PolicyTemplate, options UpdateCrossTenantAccessPolicyTemplateOperationOptions) (result UpdateCrossTenantAccessPolicyTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates",
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
