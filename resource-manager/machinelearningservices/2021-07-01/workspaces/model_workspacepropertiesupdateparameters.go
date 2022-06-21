package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacePropertiesUpdateParameters struct {
	Description                     *string                          `json:"description,omitempty"`
	FriendlyName                    *string                          `json:"friendlyName,omitempty"`
	ImageBuildCompute               *string                          `json:"imageBuildCompute,omitempty"`
	PrimaryUserAssignedIdentity     *string                          `json:"primaryUserAssignedIdentity,omitempty"`
	PublicNetworkAccess             *PublicNetworkAccess             `json:"publicNetworkAccess,omitempty"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings `json:"serviceManagedResourcesSettings,omitempty"`
}
