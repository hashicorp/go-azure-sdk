package linkers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceConfiguration struct {
	ConfigType                *LinkerConfigurationType `json:"configType,omitempty"`
	Description               *string                  `json:"description,omitempty"`
	KeyVaultReferenceIdentity *string                  `json:"keyVaultReferenceIdentity,omitempty"`
	Name                      *string                  `json:"name,omitempty"`
	Value                     *string                  `json:"value,omitempty"`
}
