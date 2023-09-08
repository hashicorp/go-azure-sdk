package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryDefinition struct {
	Discoverabilities *DirectoryDefinitionDiscoverabilities `json:"discoverabilities,omitempty"`
	DiscoveryDateTime *string                               `json:"discoveryDateTime,omitempty"`
	Id                *string                               `json:"id,omitempty"`
	Name              *string                               `json:"name,omitempty"`
	ODataType         *string                               `json:"@odata.type,omitempty"`
	Objects           *[]ObjectDefinition                   `json:"objects,omitempty"`
	ReadOnly          *bool                                 `json:"readOnly,omitempty"`
	Version           *string                               `json:"version,omitempty"`
}
