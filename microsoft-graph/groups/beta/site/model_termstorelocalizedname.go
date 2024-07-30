package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreLocalizedName struct {
	LanguageTag *string `json:"languageTag,omitempty"`
	Name        *string `json:"name,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
