package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDevicePartnerReportedThreatState string

const (
	WindowsManagedDevicePartnerReportedThreatStateactivated      WindowsManagedDevicePartnerReportedThreatState = "Activated"
	WindowsManagedDevicePartnerReportedThreatStatecompromised    WindowsManagedDevicePartnerReportedThreatState = "Compromised"
	WindowsManagedDevicePartnerReportedThreatStatedeactivated    WindowsManagedDevicePartnerReportedThreatState = "Deactivated"
	WindowsManagedDevicePartnerReportedThreatStatehighSeverity   WindowsManagedDevicePartnerReportedThreatState = "HighSeverity"
	WindowsManagedDevicePartnerReportedThreatStatelowSeverity    WindowsManagedDevicePartnerReportedThreatState = "LowSeverity"
	WindowsManagedDevicePartnerReportedThreatStatemediumSeverity WindowsManagedDevicePartnerReportedThreatState = "MediumSeverity"
	WindowsManagedDevicePartnerReportedThreatStatemisconfigured  WindowsManagedDevicePartnerReportedThreatState = "Misconfigured"
	WindowsManagedDevicePartnerReportedThreatStatesecured        WindowsManagedDevicePartnerReportedThreatState = "Secured"
	WindowsManagedDevicePartnerReportedThreatStateunknown        WindowsManagedDevicePartnerReportedThreatState = "Unknown"
	WindowsManagedDevicePartnerReportedThreatStateunresponsive   WindowsManagedDevicePartnerReportedThreatState = "Unresponsive"
)

func PossibleValuesForWindowsManagedDevicePartnerReportedThreatState() []string {
	return []string{
		string(WindowsManagedDevicePartnerReportedThreatStateactivated),
		string(WindowsManagedDevicePartnerReportedThreatStatecompromised),
		string(WindowsManagedDevicePartnerReportedThreatStatedeactivated),
		string(WindowsManagedDevicePartnerReportedThreatStatehighSeverity),
		string(WindowsManagedDevicePartnerReportedThreatStatelowSeverity),
		string(WindowsManagedDevicePartnerReportedThreatStatemediumSeverity),
		string(WindowsManagedDevicePartnerReportedThreatStatemisconfigured),
		string(WindowsManagedDevicePartnerReportedThreatStatesecured),
		string(WindowsManagedDevicePartnerReportedThreatStateunknown),
		string(WindowsManagedDevicePartnerReportedThreatStateunresponsive),
	}
}

func parseWindowsManagedDevicePartnerReportedThreatState(input string) (*WindowsManagedDevicePartnerReportedThreatState, error) {
	vals := map[string]WindowsManagedDevicePartnerReportedThreatState{
		"activated":      WindowsManagedDevicePartnerReportedThreatStateactivated,
		"compromised":    WindowsManagedDevicePartnerReportedThreatStatecompromised,
		"deactivated":    WindowsManagedDevicePartnerReportedThreatStatedeactivated,
		"highseverity":   WindowsManagedDevicePartnerReportedThreatStatehighSeverity,
		"lowseverity":    WindowsManagedDevicePartnerReportedThreatStatelowSeverity,
		"mediumseverity": WindowsManagedDevicePartnerReportedThreatStatemediumSeverity,
		"misconfigured":  WindowsManagedDevicePartnerReportedThreatStatemisconfigured,
		"secured":        WindowsManagedDevicePartnerReportedThreatStatesecured,
		"unknown":        WindowsManagedDevicePartnerReportedThreatStateunknown,
		"unresponsive":   WindowsManagedDevicePartnerReportedThreatStateunresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDevicePartnerReportedThreatState(input)
	return &out, nil
}
