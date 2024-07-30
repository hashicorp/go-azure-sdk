package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationLanguageOverride struct {
	LanguageTag         *string                                         `json:"languageTag,omitempty"`
	ODataType           *string                                         `json:"@odata.type,omitempty"`
	TranslationBehavior *TranslationLanguageOverrideTranslationBehavior `json:"translationBehavior,omitempty"`
}
