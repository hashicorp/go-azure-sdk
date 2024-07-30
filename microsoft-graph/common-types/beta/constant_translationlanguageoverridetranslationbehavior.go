package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationLanguageOverrideTranslationBehavior string

const (
	TranslationLanguageOverrideTranslationBehavior_Ask TranslationLanguageOverrideTranslationBehavior = "Ask"
	TranslationLanguageOverrideTranslationBehavior_No  TranslationLanguageOverrideTranslationBehavior = "No"
	TranslationLanguageOverrideTranslationBehavior_Yes TranslationLanguageOverrideTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationLanguageOverrideTranslationBehavior() []string {
	return []string{
		string(TranslationLanguageOverrideTranslationBehavior_Ask),
		string(TranslationLanguageOverrideTranslationBehavior_No),
		string(TranslationLanguageOverrideTranslationBehavior_Yes),
	}
}

func (s *TranslationLanguageOverrideTranslationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTranslationLanguageOverrideTranslationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTranslationLanguageOverrideTranslationBehavior(input string) (*TranslationLanguageOverrideTranslationBehavior, error) {
	vals := map[string]TranslationLanguageOverrideTranslationBehavior{
		"ask": TranslationLanguageOverrideTranslationBehavior_Ask,
		"no":  TranslationLanguageOverrideTranslationBehavior_No,
		"yes": TranslationLanguageOverrideTranslationBehavior_Yes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationLanguageOverrideTranslationBehavior(input)
	return &out, nil
}
