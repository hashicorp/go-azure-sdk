package managedhsms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmProperties struct {
	CreateMode                 *CreateMode                          `json:"createMode,omitempty"`
	EnablePurgeProtection      *bool                                `json:"enablePurgeProtection,omitempty"`
	EnableSoftDelete           *bool                                `json:"enableSoftDelete,omitempty"`
	HsmUri                     *string                              `json:"hsmUri,omitempty"`
	InitialAdminObjectIds      *[]string                            `json:"initialAdminObjectIds,omitempty"`
	NetworkAcls                *MHSMNetworkRuleSet                  `json:"networkAcls,omitempty"`
	PrivateEndpointConnections *[]MHSMPrivateEndpointConnectionItem `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState                   `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess                 `json:"publicNetworkAccess,omitempty"`
	Regions                    *[]MHSMGeoReplicatedRegion           `json:"regions,omitempty"`
	ScheduledPurgeDate         *string                              `json:"scheduledPurgeDate,omitempty"`
	SecurityDomainProperties   *ManagedHSMSecurityDomainProperties  `json:"securityDomainProperties,omitempty"`
	SoftDeleteRetentionInDays  *int64                               `json:"softDeleteRetentionInDays,omitempty"`
	StatusMessage              *string                              `json:"statusMessage,omitempty"`
	TenantId                   *string                              `json:"tenantId,omitempty"`
}
