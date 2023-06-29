package provideroperationsmetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderOperationsMetadata struct {
	DisplayName   *string              `json:"displayName,omitempty"`
	Id            *string              `json:"id,omitempty"`
	Name          *string              `json:"name,omitempty"`
	Operations    *[]ProviderOperation `json:"operations,omitempty"`
	ResourceTypes *[]ResourceType      `json:"resourceTypes,omitempty"`
	Type          *string              `json:"type,omitempty"`
}
