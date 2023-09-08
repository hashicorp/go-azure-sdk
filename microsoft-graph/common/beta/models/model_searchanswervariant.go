package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerVariant struct {
	Description *string                      `json:"description,omitempty"`
	DisplayName *string                      `json:"displayName,omitempty"`
	LanguageTag *string                      `json:"languageTag,omitempty"`
	ODataType   *string                      `json:"@odata.type,omitempty"`
	Platform    *SearchAnswerVariantPlatform `json:"platform,omitempty"`
	WebUrl      *string                      `json:"webUrl,omitempty"`
}
