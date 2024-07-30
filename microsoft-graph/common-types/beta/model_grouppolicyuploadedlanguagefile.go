package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedLanguageFile struct {
	Content              *string `json:"content,omitempty"`
	FileName             *string `json:"fileName,omitempty"`
	Id                   *string `json:"id,omitempty"`
	LanguageCode         *string `json:"languageCode,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
