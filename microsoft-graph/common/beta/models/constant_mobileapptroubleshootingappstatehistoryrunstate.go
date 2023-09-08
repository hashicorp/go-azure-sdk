package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppStateHistoryRunState string

const (
	MobileAppTroubleshootingAppStateHistoryRunStatefail          MobileAppTroubleshootingAppStateHistoryRunState = "Fail"
	MobileAppTroubleshootingAppStateHistoryRunStatenotApplicable MobileAppTroubleshootingAppStateHistoryRunState = "NotApplicable"
	MobileAppTroubleshootingAppStateHistoryRunStatepending       MobileAppTroubleshootingAppStateHistoryRunState = "Pending"
	MobileAppTroubleshootingAppStateHistoryRunStatescriptError   MobileAppTroubleshootingAppStateHistoryRunState = "ScriptError"
	MobileAppTroubleshootingAppStateHistoryRunStatesuccess       MobileAppTroubleshootingAppStateHistoryRunState = "Success"
	MobileAppTroubleshootingAppStateHistoryRunStateunknown       MobileAppTroubleshootingAppStateHistoryRunState = "Unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppStateHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppStateHistoryRunStatefail),
		string(MobileAppTroubleshootingAppStateHistoryRunStatenotApplicable),
		string(MobileAppTroubleshootingAppStateHistoryRunStatepending),
		string(MobileAppTroubleshootingAppStateHistoryRunStatescriptError),
		string(MobileAppTroubleshootingAppStateHistoryRunStatesuccess),
		string(MobileAppTroubleshootingAppStateHistoryRunStateunknown),
	}
}

func parseMobileAppTroubleshootingAppStateHistoryRunState(input string) (*MobileAppTroubleshootingAppStateHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppStateHistoryRunState{
		"fail":          MobileAppTroubleshootingAppStateHistoryRunStatefail,
		"notapplicable": MobileAppTroubleshootingAppStateHistoryRunStatenotApplicable,
		"pending":       MobileAppTroubleshootingAppStateHistoryRunStatepending,
		"scripterror":   MobileAppTroubleshootingAppStateHistoryRunStatescriptError,
		"success":       MobileAppTroubleshootingAppStateHistoryRunStatesuccess,
		"unknown":       MobileAppTroubleshootingAppStateHistoryRunStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppStateHistoryRunState(input)
	return &out, nil
}
