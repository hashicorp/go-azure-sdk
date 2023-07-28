package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceProperties struct {
	AllowPublicAccessWhenBehindVnet *bool                            `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string                          `json:"applicationInsights,omitempty"`
	AssociatedWorkspaces            *[]string                        `json:"associatedWorkspaces,omitempty"`
	ContainerRegistries             *[]string                        `json:"containerRegistries,omitempty"`
	ContainerRegistry               *string                          `json:"containerRegistry,omitempty"`
	Description                     *string                          `json:"description,omitempty"`
	DiscoveryUrl                    *string                          `json:"discoveryUrl,omitempty"`
	EnableDataIsolation             *bool                            `json:"enableDataIsolation,omitempty"`
	Encryption                      *EncryptionProperty              `json:"encryption,omitempty"`
	ExistingWorkspaces              *[]string                        `json:"existingWorkspaces,omitempty"`
	FeatureStoreSettings            *FeatureStoreSettings            `json:"featureStoreSettings,omitempty"`
	FriendlyName                    *string                          `json:"friendlyName,omitempty"`
	HbiWorkspace                    *bool                            `json:"hbiWorkspace,omitempty"`
	HubResourceId                   *string                          `json:"hubResourceId,omitempty"`
	ImageBuildCompute               *string                          `json:"imageBuildCompute,omitempty"`
	KeyVault                        *string                          `json:"keyVault,omitempty"`
	KeyVaults                       *[]string                        `json:"keyVaults,omitempty"`
	ManagedNetwork                  *ManagedNetworkSettings          `json:"managedNetwork,omitempty"`
	MlFlowTrackingUri               *string                          `json:"mlFlowTrackingUri,omitempty"`
	NotebookInfo                    *NotebookResourceInfo            `json:"notebookInfo,omitempty"`
	PrimaryUserAssignedIdentity     *string                          `json:"primaryUserAssignedIdentity,omitempty"`
	PrivateEndpointConnections      *[]PrivateEndpointConnection     `json:"privateEndpointConnections,omitempty"`
	PrivateLinkCount                *int64                           `json:"privateLinkCount,omitempty"`
	ProvisioningState               *ProvisioningState               `json:"provisioningState,omitempty"`
	PublicNetworkAccess             *PublicNetworkAccess             `json:"publicNetworkAccess,omitempty"`
	ScheduledPurgeDate              *string                          `json:"scheduledPurgeDate,omitempty"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings `json:"serviceManagedResourcesSettings,omitempty"`
	ServiceProvisionedResourceGroup *string                          `json:"serviceProvisionedResourceGroup,omitempty"`
	SharedPrivateLinkResources      *[]SharedPrivateLinkResource     `json:"sharedPrivateLinkResources,omitempty"`
	SoftDeleteRetentionInDays       *int64                           `json:"softDeleteRetentionInDays,omitempty"`
	SoftDeletedAt                   *string                          `json:"softDeletedAt,omitempty"`
	StorageAccount                  *string                          `json:"storageAccount,omitempty"`
	StorageAccounts                 *[]string                        `json:"storageAccounts,omitempty"`
	StorageHnsEnabled               *bool                            `json:"storageHnsEnabled,omitempty"`
	SystemDatastoresAuthMode        *string                          `json:"systemDatastoresAuthMode,omitempty"`
	TenantId                        *string                          `json:"tenantId,omitempty"`
	V1LegacyMode                    *bool                            `json:"v1LegacyMode,omitempty"`
	WorkspaceId                     *string                          `json:"workspaceId,omitempty"`
}
