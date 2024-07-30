package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Account struct {
	Blocked              *bool   `json:"blocked,omitempty"`
	Category             *string `json:"category,omitempty"`
	DisplayName          *string `json:"displayName,omitempty"`
	Id                   *string `json:"id,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	Number               *string `json:"number,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	SubCategory          *string `json:"subCategory,omitempty"`
}
