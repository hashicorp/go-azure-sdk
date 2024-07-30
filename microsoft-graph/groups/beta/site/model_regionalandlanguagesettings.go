package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionalAndLanguageSettings struct {
	AuthoringLanguages         *[]LocaleInfo            `json:"authoringLanguages,omitempty"`
	DefaultDisplayLanguage     *LocaleInfo              `json:"defaultDisplayLanguage,omitempty"`
	DefaultRegionalFormat      *LocaleInfo              `json:"defaultRegionalFormat,omitempty"`
	DefaultSpeechInputLanguage *LocaleInfo              `json:"defaultSpeechInputLanguage,omitempty"`
	DefaultTranslationLanguage *LocaleInfo              `json:"defaultTranslationLanguage,omitempty"`
	Id                         *string                  `json:"id,omitempty"`
	ODataType                  *string                  `json:"@odata.type,omitempty"`
	RegionalFormatOverrides    *RegionalFormatOverrides `json:"regionalFormatOverrides,omitempty"`
	TranslationPreferences     *TranslationPreferences  `json:"translationPreferences,omitempty"`
}
