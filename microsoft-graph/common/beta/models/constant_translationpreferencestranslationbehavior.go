package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationPreferencesTranslationBehavior string

const (
	TranslationPreferencesTranslationBehaviorAsk TranslationPreferencesTranslationBehavior = "Ask"
	TranslationPreferencesTranslationBehaviorNo  TranslationPreferencesTranslationBehavior = "No"
	TranslationPreferencesTranslationBehaviorYes TranslationPreferencesTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationPreferencesTranslationBehavior() []string {
	return []string{
		string(TranslationPreferencesTranslationBehaviorAsk),
		string(TranslationPreferencesTranslationBehaviorNo),
		string(TranslationPreferencesTranslationBehaviorYes),
	}
}

func parseTranslationPreferencesTranslationBehavior(input string) (*TranslationPreferencesTranslationBehavior, error) {
	vals := map[string]TranslationPreferencesTranslationBehavior{
		"ask": TranslationPreferencesTranslationBehaviorAsk,
		"no":  TranslationPreferencesTranslationBehaviorNo,
		"yes": TranslationPreferencesTranslationBehaviorYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationPreferencesTranslationBehavior(input)
	return &out, nil
}
