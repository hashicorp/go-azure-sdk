package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionLabel struct {
	Color       *string             `json:"color,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *string             `json:"id,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	Name        *string             `json:"name,omitempty"`
	ODataType   *string             `json:"@odata.type,omitempty"`
	Parent      *ParentLabelDetails `json:"parent,omitempty"`
	Sensitivity *int64              `json:"sensitivity,omitempty"`
	Tooltip     *string             `json:"tooltip,omitempty"`
}
