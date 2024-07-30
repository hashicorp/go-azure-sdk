package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRuleSignal string

const (
	WindowsUpdatesMonitoringRuleSignal_Ineligible WindowsUpdatesMonitoringRuleSignal = "ineligible"
	WindowsUpdatesMonitoringRuleSignal_Rollback   WindowsUpdatesMonitoringRuleSignal = "rollback"
)

func PossibleValuesForWindowsUpdatesMonitoringRuleSignal() []string {
	return []string{
		string(WindowsUpdatesMonitoringRuleSignal_Ineligible),
		string(WindowsUpdatesMonitoringRuleSignal_Rollback),
	}
}

func (s *WindowsUpdatesMonitoringRuleSignal) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesMonitoringRuleSignal(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesMonitoringRuleSignal(input string) (*WindowsUpdatesMonitoringRuleSignal, error) {
	vals := map[string]WindowsUpdatesMonitoringRuleSignal{
		"ineligible": WindowsUpdatesMonitoringRuleSignal_Ineligible,
		"rollback":   WindowsUpdatesMonitoringRuleSignal_Rollback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringRuleSignal(input)
	return &out, nil
}
