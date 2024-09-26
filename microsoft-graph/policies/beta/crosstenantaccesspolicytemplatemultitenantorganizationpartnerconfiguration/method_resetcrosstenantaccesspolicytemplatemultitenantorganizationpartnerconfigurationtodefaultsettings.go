package crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions() ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions {
	return ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions{}
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettings - Invoke action
// resetToDefaultSettings. Reset the cross-tenant access policy template with inbound and outbound partner configuration
// settings for a multitenant organization to the default values. In its reset state, the template has no impact on
// partner configuration settings, other than that an unconfigured partner configuration is created if needed, for every
// multitenant organization tenant.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient) ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettings(ctx context.Context, options ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationOptions) (result ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingsOperationResponse, err error) {
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
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationPartnerConfiguration/resetToDefaultSettings",
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
