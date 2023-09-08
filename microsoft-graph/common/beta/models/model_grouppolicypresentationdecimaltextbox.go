package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationDecimalTextBox struct {
	DefaultValue         *int64                 `json:"defaultValue,omitempty"`
	Definition           *GroupPolicyDefinition `json:"definition,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Label                *string                `json:"label,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	MaxValue             *int64                 `json:"maxValue,omitempty"`
	MinValue             *int64                 `json:"minValue,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	Required             *bool                  `json:"required,omitempty"`
	Spin                 *bool                  `json:"spin,omitempty"`
	SpinStep             *int64                 `json:"spinStep,omitempty"`
}
