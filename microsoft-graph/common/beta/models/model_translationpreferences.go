package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationPreferences struct {
	LanguageOverrides     *[]TranslationLanguageOverride             `json:"languageOverrides,omitempty"`
	ODataType             *string                                    `json:"@odata.type,omitempty"`
	TranslationBehavior   *TranslationPreferencesTranslationBehavior `json:"translationBehavior,omitempty"`
	UntranslatedLanguages *[]string                                  `json:"untranslatedLanguages,omitempty"`
}
