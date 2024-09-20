package crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions() ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions {
	return ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions{}
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettings - Invoke action
// resetToDefaultSettings. Reset the cross-tenant access policy template with user synchronization settings for a
// multitenant organization to the default values. In its reset state, the template has no impact on user
// synchronization settings, other than that unconfigured user synchronization settings are created if needed, for every
// multitenant organization tenant.
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient) ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettings(ctx context.Context, options ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationOptions) (result ResetCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationToDefaultSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationIdentitySynchronization/resetToDefaultSettings",
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
