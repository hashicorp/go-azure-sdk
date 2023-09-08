package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTemplate struct {
	Id                                             *string                                              `json:"id,omitempty"`
	MultiTenantOrganizationIdentitySynchronization *MultiTenantOrganizationIdentitySyncPolicyTemplate   `json:"multiTenantOrganizationIdentitySynchronization,omitempty"`
	MultiTenantOrganizationPartnerConfiguration    *MultiTenantOrganizationPartnerConfigurationTemplate `json:"multiTenantOrganizationPartnerConfiguration,omitempty"`
	ODataType                                      *string                                              `json:"@odata.type,omitempty"`
}
