package crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MultiTenantOrganizationPartnerConfigurationTemplate
}

type GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions() GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions {
	return GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions{}
}

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration - Get
// multiTenantOrganizationPartnerConfigurationTemplate. Get the cross-tenant access policy template with inbound and
// outbound partner configuration settings for a multitenant organization.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient) GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration(ctx context.Context, options GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) (result GetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationPartnerConfiguration",
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

	var model stable.MultiTenantOrganizationPartnerConfigurationTemplate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
