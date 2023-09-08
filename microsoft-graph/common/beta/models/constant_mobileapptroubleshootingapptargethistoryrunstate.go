package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppTargetHistoryRunState string

const (
	MobileAppTroubleshootingAppTargetHistoryRunStatefail          MobileAppTroubleshootingAppTargetHistoryRunState = "Fail"
	MobileAppTroubleshootingAppTargetHistoryRunStatenotApplicable MobileAppTroubleshootingAppTargetHistoryRunState = "NotApplicable"
	MobileAppTroubleshootingAppTargetHistoryRunStatepending       MobileAppTroubleshootingAppTargetHistoryRunState = "Pending"
	MobileAppTroubleshootingAppTargetHistoryRunStatescriptError   MobileAppTroubleshootingAppTargetHistoryRunState = "ScriptError"
	MobileAppTroubleshootingAppTargetHistoryRunStatesuccess       MobileAppTroubleshootingAppTargetHistoryRunState = "Success"
	MobileAppTroubleshootingAppTargetHistoryRunStateunknown       MobileAppTroubleshootingAppTargetHistoryRunState = "Unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppTargetHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppTargetHistoryRunStatefail),
		string(MobileAppTroubleshootingAppTargetHistoryRunStatenotApplicable),
		string(MobileAppTroubleshootingAppTargetHistoryRunStatepending),
		string(MobileAppTroubleshootingAppTargetHistoryRunStatescriptError),
		string(MobileAppTroubleshootingAppTargetHistoryRunStatesuccess),
		string(MobileAppTroubleshootingAppTargetHistoryRunStateunknown),
	}
}

func parseMobileAppTroubleshootingAppTargetHistoryRunState(input string) (*MobileAppTroubleshootingAppTargetHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppTargetHistoryRunState{
		"fail":          MobileAppTroubleshootingAppTargetHistoryRunStatefail,
		"notapplicable": MobileAppTroubleshootingAppTargetHistoryRunStatenotApplicable,
		"pending":       MobileAppTroubleshootingAppTargetHistoryRunStatepending,
		"scripterror":   MobileAppTroubleshootingAppTargetHistoryRunStatescriptError,
		"success":       MobileAppTroubleshootingAppTargetHistoryRunStatesuccess,
		"unknown":       MobileAppTroubleshootingAppTargetHistoryRunStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppTargetHistoryRunState(input)
	return &out, nil
}
