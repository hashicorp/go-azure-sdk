package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingSource struct {
	Expression *string                                     `json:"expression,omitempty"`
	Name       *string                                     `json:"name,omitempty"`
	ODataType  *string                                     `json:"@odata.type,omitempty"`
	Parameters *[]StringKeyAttributeMappingSourceValuePair `json:"parameters,omitempty"`
	Type       *AttributeMappingSourceType                 `json:"type,omitempty"`
}
