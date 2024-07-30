package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySensitivityLabel struct {
	Color          *string                   `json:"color,omitempty"`
	ContentFormats *[]string                 `json:"contentFormats,omitempty"`
	Description    *string                   `json:"description,omitempty"`
	HasProtection  *bool                     `json:"hasProtection,omitempty"`
	Id             *string                   `json:"id,omitempty"`
	IsActive       *bool                     `json:"isActive,omitempty"`
	IsAppliable    *bool                     `json:"isAppliable,omitempty"`
	Name           *string                   `json:"name,omitempty"`
	ODataType      *string                   `json:"@odata.type,omitempty"`
	Parent         *SecuritySensitivityLabel `json:"parent,omitempty"`
	Sensitivity    *int64                    `json:"sensitivity,omitempty"`
	Tooltip        *string                   `json:"tooltip,omitempty"`
}
