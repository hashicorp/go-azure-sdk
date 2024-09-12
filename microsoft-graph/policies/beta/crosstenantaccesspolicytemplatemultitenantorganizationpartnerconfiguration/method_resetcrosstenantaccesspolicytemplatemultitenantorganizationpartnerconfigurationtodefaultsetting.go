package crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSetting - Invoke action
// resetToDefaultSettings. Reset the cross-tenant access policy template with inbound and outbound partner configuration
// settings for a multitenant organization to the default values. In its reset state, the template has no impact on
// partner configuration settings, other than that an unconfigured partner configuration is created if needed, for every
// multitenant organization tenant.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient) ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSetting(ctx context.Context) (result ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationToDefaultSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationPartnerConfiguration/resetToDefaultSettings",
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
