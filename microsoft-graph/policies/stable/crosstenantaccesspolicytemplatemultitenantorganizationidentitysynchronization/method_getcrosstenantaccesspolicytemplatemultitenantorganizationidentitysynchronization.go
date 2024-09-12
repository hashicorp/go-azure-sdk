package crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MultiTenantOrganizationIdentitySyncPolicyTemplate
}

type GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions() GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions {
	return GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions{}
}

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization - Get
// multiTenantOrganizationIdentitySyncPolicyTemplate. Get the cross-tenant access policy template with user
// synchronization settings for a multitenant organization.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient) GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization(ctx context.Context, options GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) (result GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationIdentitySynchronization",
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

	var model stable.MultiTenantOrganizationIdentitySyncPolicyTemplate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
