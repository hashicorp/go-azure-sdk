package addons

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArcAddonProperties struct {
	HostPlatform      *PlatformType     `json:"hostPlatform,omitempty"`
	HostPlatformType  *HostPlatformType `json:"hostPlatformType,omitempty"`
	ProvisioningState *AddonState       `json:"provisioningState,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName"`
	ResourceLocation  string            `json:"resourceLocation"`
	ResourceName      string            `json:"resourceName"`
	SubscriptionId    string            `json:"subscriptionId"`
	Version           *string           `json:"version,omitempty"`
}
