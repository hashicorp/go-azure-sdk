package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationPreferencesTranslationBehavior string

const (
	TranslationPreferencesTranslationBehavior_Ask TranslationPreferencesTranslationBehavior = "Ask"
	TranslationPreferencesTranslationBehavior_No  TranslationPreferencesTranslationBehavior = "No"
	TranslationPreferencesTranslationBehavior_Yes TranslationPreferencesTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationPreferencesTranslationBehavior() []string {
	return []string{
		string(TranslationPreferencesTranslationBehavior_Ask),
		string(TranslationPreferencesTranslationBehavior_No),
		string(TranslationPreferencesTranslationBehavior_Yes),
	}
}

func (s *TranslationPreferencesTranslationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTranslationPreferencesTranslationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTranslationPreferencesTranslationBehavior(input string) (*TranslationPreferencesTranslationBehavior, error) {
	vals := map[string]TranslationPreferencesTranslationBehavior{
		"ask": TranslationPreferencesTranslationBehavior_Ask,
		"no":  TranslationPreferencesTranslationBehavior_No,
		"yes": TranslationPreferencesTranslationBehavior_Yes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationPreferencesTranslationBehavior(input)
	return &out, nil
}
