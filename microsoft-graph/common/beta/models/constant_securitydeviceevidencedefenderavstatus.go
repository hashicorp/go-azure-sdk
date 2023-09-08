package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceDefenderAvStatus string

const (
	SecurityDeviceEvidenceDefenderAvStatusdisabled     SecurityDeviceEvidenceDefenderAvStatus = "Disabled"
	SecurityDeviceEvidenceDefenderAvStatusnotReporting SecurityDeviceEvidenceDefenderAvStatus = "NotReporting"
	SecurityDeviceEvidenceDefenderAvStatusnotSupported SecurityDeviceEvidenceDefenderAvStatus = "NotSupported"
	SecurityDeviceEvidenceDefenderAvStatusnotUpdated   SecurityDeviceEvidenceDefenderAvStatus = "NotUpdated"
	SecurityDeviceEvidenceDefenderAvStatusunknown      SecurityDeviceEvidenceDefenderAvStatus = "Unknown"
	SecurityDeviceEvidenceDefenderAvStatusupdated      SecurityDeviceEvidenceDefenderAvStatus = "Updated"
)

func PossibleValuesForSecurityDeviceEvidenceDefenderAvStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceDefenderAvStatusdisabled),
		string(SecurityDeviceEvidenceDefenderAvStatusnotReporting),
		string(SecurityDeviceEvidenceDefenderAvStatusnotSupported),
		string(SecurityDeviceEvidenceDefenderAvStatusnotUpdated),
		string(SecurityDeviceEvidenceDefenderAvStatusunknown),
		string(SecurityDeviceEvidenceDefenderAvStatusupdated),
	}
}

func parseSecurityDeviceEvidenceDefenderAvStatus(input string) (*SecurityDeviceEvidenceDefenderAvStatus, error) {
	vals := map[string]SecurityDeviceEvidenceDefenderAvStatus{
		"disabled":     SecurityDeviceEvidenceDefenderAvStatusdisabled,
		"notreporting": SecurityDeviceEvidenceDefenderAvStatusnotReporting,
		"notsupported": SecurityDeviceEvidenceDefenderAvStatusnotSupported,
		"notupdated":   SecurityDeviceEvidenceDefenderAvStatusnotUpdated,
		"unknown":      SecurityDeviceEvidenceDefenderAvStatusunknown,
		"updated":      SecurityDeviceEvidenceDefenderAvStatusupdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceDefenderAvStatus(input)
	return &out, nil
}
