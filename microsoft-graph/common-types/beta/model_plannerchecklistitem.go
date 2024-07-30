package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerChecklistItem struct {
	IsChecked            *bool        `json:"isChecked,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	OrderHint            *string      `json:"orderHint,omitempty"`
	Title                *string      `json:"title,omitempty"`
}
