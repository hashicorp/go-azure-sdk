package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDevicePartnerReportedThreatState string

const (
	WindowsManagedDevicePartnerReportedThreatState_Activated      WindowsManagedDevicePartnerReportedThreatState = "activated"
	WindowsManagedDevicePartnerReportedThreatState_Compromised    WindowsManagedDevicePartnerReportedThreatState = "compromised"
	WindowsManagedDevicePartnerReportedThreatState_Deactivated    WindowsManagedDevicePartnerReportedThreatState = "deactivated"
	WindowsManagedDevicePartnerReportedThreatState_HighSeverity   WindowsManagedDevicePartnerReportedThreatState = "highSeverity"
	WindowsManagedDevicePartnerReportedThreatState_LowSeverity    WindowsManagedDevicePartnerReportedThreatState = "lowSeverity"
	WindowsManagedDevicePartnerReportedThreatState_MediumSeverity WindowsManagedDevicePartnerReportedThreatState = "mediumSeverity"
	WindowsManagedDevicePartnerReportedThreatState_Misconfigured  WindowsManagedDevicePartnerReportedThreatState = "misconfigured"
	WindowsManagedDevicePartnerReportedThreatState_Secured        WindowsManagedDevicePartnerReportedThreatState = "secured"
	WindowsManagedDevicePartnerReportedThreatState_Unknown        WindowsManagedDevicePartnerReportedThreatState = "unknown"
	WindowsManagedDevicePartnerReportedThreatState_Unresponsive   WindowsManagedDevicePartnerReportedThreatState = "unresponsive"
)

func PossibleValuesForWindowsManagedDevicePartnerReportedThreatState() []string {
	return []string{
		string(WindowsManagedDevicePartnerReportedThreatState_Activated),
		string(WindowsManagedDevicePartnerReportedThreatState_Compromised),
		string(WindowsManagedDevicePartnerReportedThreatState_Deactivated),
		string(WindowsManagedDevicePartnerReportedThreatState_HighSeverity),
		string(WindowsManagedDevicePartnerReportedThreatState_LowSeverity),
		string(WindowsManagedDevicePartnerReportedThreatState_MediumSeverity),
		string(WindowsManagedDevicePartnerReportedThreatState_Misconfigured),
		string(WindowsManagedDevicePartnerReportedThreatState_Secured),
		string(WindowsManagedDevicePartnerReportedThreatState_Unknown),
		string(WindowsManagedDevicePartnerReportedThreatState_Unresponsive),
	}
}

func (s *WindowsManagedDevicePartnerReportedThreatState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDevicePartnerReportedThreatState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDevicePartnerReportedThreatState(input string) (*WindowsManagedDevicePartnerReportedThreatState, error) {
	vals := map[string]WindowsManagedDevicePartnerReportedThreatState{
		"activated":      WindowsManagedDevicePartnerReportedThreatState_Activated,
		"compromised":    WindowsManagedDevicePartnerReportedThreatState_Compromised,
		"deactivated":    WindowsManagedDevicePartnerReportedThreatState_Deactivated,
		"highseverity":   WindowsManagedDevicePartnerReportedThreatState_HighSeverity,
		"lowseverity":    WindowsManagedDevicePartnerReportedThreatState_LowSeverity,
		"mediumseverity": WindowsManagedDevicePartnerReportedThreatState_MediumSeverity,
		"misconfigured":  WindowsManagedDevicePartnerReportedThreatState_Misconfigured,
		"secured":        WindowsManagedDevicePartnerReportedThreatState_Secured,
		"unknown":        WindowsManagedDevicePartnerReportedThreatState_Unknown,
		"unresponsive":   WindowsManagedDevicePartnerReportedThreatState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDevicePartnerReportedThreatState(input)
	return &out, nil
}
