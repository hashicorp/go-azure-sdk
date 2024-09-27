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

type UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions() UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions {
	return UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions{}
}

func (o UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration - Update
// multiTenantOrganizationPartnerConfigurationTemplate. Update the cross-tenant access policy template with inbound and
// outbound partner configuration settings for a multitenant organization.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient) UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration(ctx context.Context, input stable.MultiTenantOrganizationPartnerConfigurationTemplate, options UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationOptions) (result UpdateCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationPartnerConfiguration",
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
