package privatelinkforazuread

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkPolicy struct {
	AllTenants     *bool              `json:"allTenants,omitempty"`
	Id             *string            `json:"id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	OwnerTenantId  *string            `json:"ownerTenantId,omitempty"`
	ResourceGroup  *string            `json:"resourceGroup,omitempty"`
	ResourceName   *string            `json:"resourceName,omitempty"`
	SubscriptionId *string            `json:"subscriptionId,omitempty"`
	Tags           *map[string]string `json:"tags,omitempty"`
	Tenants        *[]string          `json:"tenants,omitempty"`
	Type           *string            `json:"type,omitempty"`
}
