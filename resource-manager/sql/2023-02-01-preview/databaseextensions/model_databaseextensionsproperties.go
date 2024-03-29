package databaseextensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseExtensionsProperties struct {
	AdministratorLogin         *string                   `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                   `json:"administratorLoginPassword,omitempty"`
	AuthenticationType         *string                   `json:"authenticationType,omitempty"`
	DatabaseEdition            *string                   `json:"databaseEdition,omitempty"`
	MaxSizeBytes               *string                   `json:"maxSizeBytes,omitempty"`
	NetworkIsolation           *NetworkIsolationSettings `json:"networkIsolation,omitempty"`
	OperationMode              OperationMode             `json:"operationMode"`
	ServiceObjectiveName       *string                   `json:"serviceObjectiveName,omitempty"`
	StorageKey                 string                    `json:"storageKey"`
	StorageKeyType             StorageKeyType            `json:"storageKeyType"`
	StorageUri                 string                    `json:"storageUri"`
}
