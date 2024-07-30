package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantRelationship struct {
	DelegatedAdminCustomers     *[]DelegatedAdminCustomer     `json:"delegatedAdminCustomers,omitempty"`
	DelegatedAdminRelationships *[]DelegatedAdminRelationship `json:"delegatedAdminRelationships,omitempty"`
	ManagedTenants              *ManagedTenantsManagedTenant  `json:"managedTenants,omitempty"`
	MultiTenantOrganization     *MultiTenantOrganization      `json:"multiTenantOrganization,omitempty"`
	ODataType                   *string                       `json:"@odata.type,omitempty"`
}
