package managedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterIdentity struct {
	DelegatedResources     *map[string]DelegatedResource                `json:"delegatedResources,omitempty"`
	PrincipalId            *string                                      `json:"principalId,omitempty"`
	TenantId               *string                                      `json:"tenantId,omitempty"`
	Type                   *ResourceIdentityType                        `json:"type,omitempty"`
	UserAssignedIdentities *map[string]UserAssignedIdentitiesProperties `json:"userAssignedIdentities,omitempty"`
}
