package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataReferenceDefinition struct {
	Code                 *string `json:"code,omitempty"`
	CreatedDateTime      *string `json:"createdDateTime,omitempty"`
	Id                   *string `json:"id,omitempty"`
	IsDisabled           *bool   `json:"isDisabled,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ReferenceType        *string `json:"referenceType,omitempty"`
	SortIndex            *int64  `json:"sortIndex,omitempty"`
	Source               *string `json:"source,omitempty"`
}
