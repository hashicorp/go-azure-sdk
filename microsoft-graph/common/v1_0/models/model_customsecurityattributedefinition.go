package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeDefinition struct {
	AllowedValues           *[]AllowedValue `json:"allowedValues,omitempty"`
	AttributeSet            *string         `json:"attributeSet,omitempty"`
	Description             *string         `json:"description,omitempty"`
	Id                      *string         `json:"id,omitempty"`
	IsCollection            *bool           `json:"isCollection,omitempty"`
	IsSearchable            *bool           `json:"isSearchable,omitempty"`
	Name                    *string         `json:"name,omitempty"`
	ODataType               *string         `json:"@odata.type,omitempty"`
	Status                  *string         `json:"status,omitempty"`
	Type                    *string         `json:"type,omitempty"`
	UsePreDefinedValuesOnly *bool           `json:"usePreDefinedValuesOnly,omitempty"`
}
