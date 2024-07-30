package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreStore struct {
	DefaultLanguageTag *string           `json:"defaultLanguageTag,omitempty"`
	Groups             *[]TermStoreGroup `json:"groups,omitempty"`
	Id                 *string           `json:"id,omitempty"`
	LanguageTags       *[]string         `json:"languageTags,omitempty"`
	ODataType          *string           `json:"@odata.type,omitempty"`
	Sets               *[]TermStoreSet   `json:"sets,omitempty"`
}
