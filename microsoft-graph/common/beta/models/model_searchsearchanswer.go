package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchSearchAnswer struct {
	Description          *string            `json:"description,omitempty"`
	DisplayName          *string            `json:"displayName,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	LastModifiedBy       *SearchIdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string            `json:"@odata.type,omitempty"`
	WebUrl               *string            `json:"webUrl,omitempty"`
}
