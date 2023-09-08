package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UploadSession struct {
	ExpirationDateTime *string   `json:"expirationDateTime,omitempty"`
	NextExpectedRanges *[]string `json:"nextExpectedRanges,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
	UploadUrl          *string   `json:"uploadUrl,omitempty"`
}
