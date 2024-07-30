package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingLanguageDetail struct {
	Content              *string        `json:"content,omitempty"`
	CreatedBy            *EmailIdentity `json:"createdBy,omitempty"`
	CreatedDateTime      *string        `json:"createdDateTime,omitempty"`
	Description          *string        `json:"description,omitempty"`
	DisplayName          *string        `json:"displayName,omitempty"`
	Id                   *string        `json:"id,omitempty"`
	IsDefaultLangauge    *bool          `json:"isDefaultLangauge,omitempty"`
	LastModifiedBy       *EmailIdentity `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string        `json:"lastModifiedDateTime,omitempty"`
	Locale               *string        `json:"locale,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
}
