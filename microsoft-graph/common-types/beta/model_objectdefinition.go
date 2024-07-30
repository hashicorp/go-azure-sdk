package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDefinition struct {
	Attributes    *[]AttributeDefinition           `json:"attributes,omitempty"`
	Metadata      *[]ObjectDefinitionMetadataEntry `json:"metadata,omitempty"`
	Name          *string                          `json:"name,omitempty"`
	ODataType     *string                          `json:"@odata.type,omitempty"`
	SupportedApis *[]string                        `json:"supportedApis,omitempty"`
}
