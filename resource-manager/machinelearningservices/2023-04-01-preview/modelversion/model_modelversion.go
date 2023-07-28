package modelversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelVersion struct {
	AutoDeleteSetting    *AutoDeleteSetting      `json:"autoDeleteSetting,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	Flavors              *map[string]FlavorData  `json:"flavors,omitempty"`
	IntellectualProperty *IntellectualProperty   `json:"intellectualProperty,omitempty"`
	IsAnonymous          *bool                   `json:"isAnonymous,omitempty"`
	IsArchived           *bool                   `json:"isArchived,omitempty"`
	JobName              *string                 `json:"jobName,omitempty"`
	ModelType            *string                 `json:"modelType,omitempty"`
	ModelUri             *string                 `json:"modelUri,omitempty"`
	Properties           *map[string]string      `json:"properties,omitempty"`
	ProvisioningState    *AssetProvisioningState `json:"provisioningState,omitempty"`
	Stage                *string                 `json:"stage,omitempty"`
	Tags                 *map[string]string      `json:"tags,omitempty"`
}
