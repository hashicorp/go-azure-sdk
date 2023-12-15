package provideroperationsmetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceType struct {
	DisplayName *string              `json:"displayName,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Operations  *[]ProviderOperation `json:"operations,omitempty"`
}
