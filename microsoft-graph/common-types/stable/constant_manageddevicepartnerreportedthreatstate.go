package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDevicePartnerReportedThreatState string

const (
	ManagedDevicePartnerReportedThreatState_Activated      ManagedDevicePartnerReportedThreatState = "activated"
	ManagedDevicePartnerReportedThreatState_Compromised    ManagedDevicePartnerReportedThreatState = "compromised"
	ManagedDevicePartnerReportedThreatState_Deactivated    ManagedDevicePartnerReportedThreatState = "deactivated"
	ManagedDevicePartnerReportedThreatState_HighSeverity   ManagedDevicePartnerReportedThreatState = "highSeverity"
	ManagedDevicePartnerReportedThreatState_LowSeverity    ManagedDevicePartnerReportedThreatState = "lowSeverity"
	ManagedDevicePartnerReportedThreatState_MediumSeverity ManagedDevicePartnerReportedThreatState = "mediumSeverity"
	ManagedDevicePartnerReportedThreatState_Misconfigured  ManagedDevicePartnerReportedThreatState = "misconfigured"
	ManagedDevicePartnerReportedThreatState_Secured        ManagedDevicePartnerReportedThreatState = "secured"
	ManagedDevicePartnerReportedThreatState_Unknown        ManagedDevicePartnerReportedThreatState = "unknown"
	ManagedDevicePartnerReportedThreatState_Unresponsive   ManagedDevicePartnerReportedThreatState = "unresponsive"
)

func PossibleValuesForManagedDevicePartnerReportedThreatState() []string {
	return []string{
		string(ManagedDevicePartnerReportedThreatState_Activated),
		string(ManagedDevicePartnerReportedThreatState_Compromised),
		string(ManagedDevicePartnerReportedThreatState_Deactivated),
		string(ManagedDevicePartnerReportedThreatState_HighSeverity),
		string(ManagedDevicePartnerReportedThreatState_LowSeverity),
		string(ManagedDevicePartnerReportedThreatState_MediumSeverity),
		string(ManagedDevicePartnerReportedThreatState_Misconfigured),
		string(ManagedDevicePartnerReportedThreatState_Secured),
		string(ManagedDevicePartnerReportedThreatState_Unknown),
		string(ManagedDevicePartnerReportedThreatState_Unresponsive),
	}
}

func (s *ManagedDevicePartnerReportedThreatState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDevicePartnerReportedThreatState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDevicePartnerReportedThreatState(input string) (*ManagedDevicePartnerReportedThreatState, error) {
	vals := map[string]ManagedDevicePartnerReportedThreatState{
		"activated":      ManagedDevicePartnerReportedThreatState_Activated,
		"compromised":    ManagedDevicePartnerReportedThreatState_Compromised,
		"deactivated":    ManagedDevicePartnerReportedThreatState_Deactivated,
		"highseverity":   ManagedDevicePartnerReportedThreatState_HighSeverity,
		"lowseverity":    ManagedDevicePartnerReportedThreatState_LowSeverity,
		"mediumseverity": ManagedDevicePartnerReportedThreatState_MediumSeverity,
		"misconfigured":  ManagedDevicePartnerReportedThreatState_Misconfigured,
		"secured":        ManagedDevicePartnerReportedThreatState_Secured,
		"unknown":        ManagedDevicePartnerReportedThreatState_Unknown,
		"unresponsive":   ManagedDevicePartnerReportedThreatState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDevicePartnerReportedThreatState(input)
	return &out, nil
}
