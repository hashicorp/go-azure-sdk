package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingParameterSchema struct {
	AllowMultipleOccurrences *bool                                `json:"allowMultipleOccurrences,omitempty"`
	Name                     *string                              `json:"name,omitempty"`
	ODataType                *string                              `json:"@odata.type,omitempty"`
	Required                 *bool                                `json:"required,omitempty"`
	Type                     *AttributeMappingParameterSchemaType `json:"type,omitempty"`
}
