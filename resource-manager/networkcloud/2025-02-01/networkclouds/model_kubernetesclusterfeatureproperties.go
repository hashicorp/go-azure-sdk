package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterFeatureProperties struct {
	AvailabilityLifecycle *KubernetesClusterFeatureAvailabilityLifecycle `json:"availabilityLifecycle,omitempty"`
	DetailedStatus        *KubernetesClusterFeatureDetailedStatus        `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                                        `json:"detailedStatusMessage,omitempty"`
	Options               *[]StringKeyValuePair                          `json:"options,omitempty"`
	ProvisioningState     *KubernetesClusterFeatureProvisioningState     `json:"provisioningState,omitempty"`
	Required              *KubernetesClusterFeatureRequired              `json:"required,omitempty"`
	Version               *string                                        `json:"version,omitempty"`
}
