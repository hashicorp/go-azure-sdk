package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportSuspiciousActivitySettingsState string

const (
	ReportSuspiciousActivitySettingsState_Default  ReportSuspiciousActivitySettingsState = "default"
	ReportSuspiciousActivitySettingsState_Disabled ReportSuspiciousActivitySettingsState = "disabled"
	ReportSuspiciousActivitySettingsState_Enabled  ReportSuspiciousActivitySettingsState = "enabled"
)

func PossibleValuesForReportSuspiciousActivitySettingsState() []string {
	return []string{
		string(ReportSuspiciousActivitySettingsState_Default),
		string(ReportSuspiciousActivitySettingsState_Disabled),
		string(ReportSuspiciousActivitySettingsState_Enabled),
	}
}

func (s *ReportSuspiciousActivitySettingsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReportSuspiciousActivitySettingsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReportSuspiciousActivitySettingsState(input string) (*ReportSuspiciousActivitySettingsState, error) {
	vals := map[string]ReportSuspiciousActivitySettingsState{
		"default":  ReportSuspiciousActivitySettingsState_Default,
		"disabled": ReportSuspiciousActivitySettingsState_Disabled,
		"enabled":  ReportSuspiciousActivitySettingsState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReportSuspiciousActivitySettingsState(input)
	return &out, nil
}
