package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationTextBox struct {
	DefaultValue         *string                `json:"defaultValue,omitempty"`
	Definition           *GroupPolicyDefinition `json:"definition,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Label                *string                `json:"label,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	MaxLength            *int64                 `json:"maxLength,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	Required             *bool                  `json:"required,omitempty"`
}
