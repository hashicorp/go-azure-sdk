package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentExtensionConfigItem struct {
	KeyVaultReference *KeyVaultParameterReference  `json:"keyVaultReference,omitempty"`
	Type              *ExtensionConfigPropertyType `json:"type,omitempty"`
	Value             *interface{}                 `json:"value,omitempty"`
}
