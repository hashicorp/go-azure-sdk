package instance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstanceProperties struct {
	AdrNamespaceRef               *AzureDeviceRegistryNamespaceRef `json:"adrNamespaceRef,omitempty"`
	DefaultSecretProviderClassRef *SecretProviderClassRef          `json:"defaultSecretProviderClassRef,omitempty"`
	Description                   *string                          `json:"description,omitempty"`
	Features                      *map[string]InstanceFeature      `json:"features,omitempty"`
	HealthState                   *ResourceHealthState             `json:"healthState,omitempty"`
	ProvisioningState             *ProvisioningState               `json:"provisioningState,omitempty"`
	SchemaRegistryRef             SchemaRegistryRef                `json:"schemaRegistryRef"`
	Version                       *string                          `json:"version,omitempty"`
}
