package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationLanguageOverrideTranslationBehavior string

const (
	TranslationLanguageOverrideTranslationBehaviorAsk TranslationLanguageOverrideTranslationBehavior = "Ask"
	TranslationLanguageOverrideTranslationBehaviorNo  TranslationLanguageOverrideTranslationBehavior = "No"
	TranslationLanguageOverrideTranslationBehaviorYes TranslationLanguageOverrideTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationLanguageOverrideTranslationBehavior() []string {
	return []string{
		string(TranslationLanguageOverrideTranslationBehaviorAsk),
		string(TranslationLanguageOverrideTranslationBehaviorNo),
		string(TranslationLanguageOverrideTranslationBehaviorYes),
	}
}

func parseTranslationLanguageOverrideTranslationBehavior(input string) (*TranslationLanguageOverrideTranslationBehavior, error) {
	vals := map[string]TranslationLanguageOverrideTranslationBehavior{
		"ask": TranslationLanguageOverrideTranslationBehaviorAsk,
		"no":  TranslationLanguageOverrideTranslationBehaviorNo,
		"yes": TranslationLanguageOverrideTranslationBehaviorYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationLanguageOverrideTranslationBehavior(input)
	return &out, nil
}
