package featurestoreentityversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturestoreEntityVersion struct {
	Description       *string                 `json:"description,omitempty"`
	IndexColumns      *[]IndexColumn          `json:"indexColumns,omitempty"`
	IsAnonymous       *bool                   `json:"isAnonymous,omitempty"`
	IsArchived        *bool                   `json:"isArchived,omitempty"`
	Properties        *map[string]string      `json:"properties,omitempty"`
	ProvisioningState *AssetProvisioningState `json:"provisioningState,omitempty"`
	Stage             *string                 `json:"stage,omitempty"`
	Tags              *map[string]string      `json:"tags,omitempty"`
}
