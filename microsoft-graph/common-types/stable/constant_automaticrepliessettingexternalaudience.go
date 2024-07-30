package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSettingExternalAudience string

const (
	AutomaticRepliesSettingExternalAudience_All          AutomaticRepliesSettingExternalAudience = "all"
	AutomaticRepliesSettingExternalAudience_ContactsOnly AutomaticRepliesSettingExternalAudience = "contactsOnly"
	AutomaticRepliesSettingExternalAudience_None         AutomaticRepliesSettingExternalAudience = "none"
)

func PossibleValuesForAutomaticRepliesSettingExternalAudience() []string {
	return []string{
		string(AutomaticRepliesSettingExternalAudience_All),
		string(AutomaticRepliesSettingExternalAudience_ContactsOnly),
		string(AutomaticRepliesSettingExternalAudience_None),
	}
}

func (s *AutomaticRepliesSettingExternalAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticRepliesSettingExternalAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticRepliesSettingExternalAudience(input string) (*AutomaticRepliesSettingExternalAudience, error) {
	vals := map[string]AutomaticRepliesSettingExternalAudience{
		"all":          AutomaticRepliesSettingExternalAudience_All,
		"contactsonly": AutomaticRepliesSettingExternalAudience_ContactsOnly,
		"none":         AutomaticRepliesSettingExternalAudience_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingExternalAudience(input)
	return &out, nil
}
