package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesRoleProperties struct {
	HostPlatform            PlatformType            `json:"hostPlatform"`
	HostPlatformType        *HostPlatformType       `json:"hostPlatformType,omitempty"`
	KubernetesClusterInfo   KubernetesClusterInfo   `json:"kubernetesClusterInfo"`
	KubernetesRoleResources KubernetesRoleResources `json:"kubernetesRoleResources"`
	ProvisioningState       *KubernetesState        `json:"provisioningState,omitempty"`
	RoleStatus              RoleStatus              `json:"roleStatus"`
}
