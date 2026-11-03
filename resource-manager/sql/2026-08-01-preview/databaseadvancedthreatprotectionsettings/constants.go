package databaseadvancedthreatprotectionsettings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionState string

const (
	AdvancedThreatProtectionStateDisabled AdvancedThreatProtectionState = "Disabled"
	AdvancedThreatProtectionStateEnabled  AdvancedThreatProtectionState = "Enabled"
	AdvancedThreatProtectionStateNew      AdvancedThreatProtectionState = "New"
)

func PossibleValuesForAdvancedThreatProtectionState() []string {
	return []string{
		string(AdvancedThreatProtectionStateDisabled),
		string(AdvancedThreatProtectionStateEnabled),
		string(AdvancedThreatProtectionStateNew),
	}
}

func (s *AdvancedThreatProtectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvancedThreatProtectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvancedThreatProtectionState(input string) (*AdvancedThreatProtectionState, error) {
	vals := map[string]AdvancedThreatProtectionState{
		"disabled": AdvancedThreatProtectionStateDisabled,
		"enabled":  AdvancedThreatProtectionStateEnabled,
		"new":      AdvancedThreatProtectionStateNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedThreatProtectionState(input)
	return &out, nil
}
