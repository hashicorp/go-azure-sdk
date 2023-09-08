package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDevicePartnerReportedThreatState string

const (
	ManagedDevicePartnerReportedThreatStateactivated      ManagedDevicePartnerReportedThreatState = "Activated"
	ManagedDevicePartnerReportedThreatStatecompromised    ManagedDevicePartnerReportedThreatState = "Compromised"
	ManagedDevicePartnerReportedThreatStatedeactivated    ManagedDevicePartnerReportedThreatState = "Deactivated"
	ManagedDevicePartnerReportedThreatStatehighSeverity   ManagedDevicePartnerReportedThreatState = "HighSeverity"
	ManagedDevicePartnerReportedThreatStatelowSeverity    ManagedDevicePartnerReportedThreatState = "LowSeverity"
	ManagedDevicePartnerReportedThreatStatemediumSeverity ManagedDevicePartnerReportedThreatState = "MediumSeverity"
	ManagedDevicePartnerReportedThreatStatemisconfigured  ManagedDevicePartnerReportedThreatState = "Misconfigured"
	ManagedDevicePartnerReportedThreatStatesecured        ManagedDevicePartnerReportedThreatState = "Secured"
	ManagedDevicePartnerReportedThreatStateunknown        ManagedDevicePartnerReportedThreatState = "Unknown"
	ManagedDevicePartnerReportedThreatStateunresponsive   ManagedDevicePartnerReportedThreatState = "Unresponsive"
)

func PossibleValuesForManagedDevicePartnerReportedThreatState() []string {
	return []string{
		string(ManagedDevicePartnerReportedThreatStateactivated),
		string(ManagedDevicePartnerReportedThreatStatecompromised),
		string(ManagedDevicePartnerReportedThreatStatedeactivated),
		string(ManagedDevicePartnerReportedThreatStatehighSeverity),
		string(ManagedDevicePartnerReportedThreatStatelowSeverity),
		string(ManagedDevicePartnerReportedThreatStatemediumSeverity),
		string(ManagedDevicePartnerReportedThreatStatemisconfigured),
		string(ManagedDevicePartnerReportedThreatStatesecured),
		string(ManagedDevicePartnerReportedThreatStateunknown),
		string(ManagedDevicePartnerReportedThreatStateunresponsive),
	}
}

func parseManagedDevicePartnerReportedThreatState(input string) (*ManagedDevicePartnerReportedThreatState, error) {
	vals := map[string]ManagedDevicePartnerReportedThreatState{
		"activated":      ManagedDevicePartnerReportedThreatStateactivated,
		"compromised":    ManagedDevicePartnerReportedThreatStatecompromised,
		"deactivated":    ManagedDevicePartnerReportedThreatStatedeactivated,
		"highseverity":   ManagedDevicePartnerReportedThreatStatehighSeverity,
		"lowseverity":    ManagedDevicePartnerReportedThreatStatelowSeverity,
		"mediumseverity": ManagedDevicePartnerReportedThreatStatemediumSeverity,
		"misconfigured":  ManagedDevicePartnerReportedThreatStatemisconfigured,
		"secured":        ManagedDevicePartnerReportedThreatStatesecured,
		"unknown":        ManagedDevicePartnerReportedThreatStateunknown,
		"unresponsive":   ManagedDevicePartnerReportedThreatStateunresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDevicePartnerReportedThreatState(input)
	return &out, nil
}
