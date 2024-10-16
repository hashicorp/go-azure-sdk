package advancedthreatprotectionsettings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatProtectionState string

const (
	ThreatProtectionStateDisabled ThreatProtectionState = "Disabled"
	ThreatProtectionStateEnabled  ThreatProtectionState = "Enabled"
)

func PossibleValuesForThreatProtectionState() []string {
	return []string{
		string(ThreatProtectionStateDisabled),
		string(ThreatProtectionStateEnabled),
	}
}

func (s *ThreatProtectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatProtectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatProtectionState(input string) (*ThreatProtectionState, error) {
	vals := map[string]ThreatProtectionState{
		"disabled": ThreatProtectionStateDisabled,
		"enabled":  ThreatProtectionStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatProtectionState(input)
	return &out, nil
}
