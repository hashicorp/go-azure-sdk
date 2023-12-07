package settings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingsKind string

const (
	SettingsKindTaginheritance SettingsKind = "taginheritance"
)

func PossibleValuesForSettingsKind() []string {
	return []string{
		string(SettingsKindTaginheritance),
	}
}

func (s *SettingsKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingsKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSettingsKind(input string) (*SettingsKind, error) {
	vals := map[string]SettingsKind{
		"taginheritance": SettingsKindTaginheritance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingsKind(input)
	return &out, nil
}
