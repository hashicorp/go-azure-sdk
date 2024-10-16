package schemaregistries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaRegistryProperties struct {
	Description                *string            `json:"description,omitempty"`
	DisplayName                *string            `json:"displayName,omitempty"`
	Namespace                  string             `json:"namespace"`
	ProvisioningState          *ProvisioningState `json:"provisioningState,omitempty"`
	StorageAccountContainerURL string             `json:"storageAccountContainerUrl"`
	Uuid                       *string            `json:"uuid,omitempty"`
}
