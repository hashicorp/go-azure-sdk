package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppPolicyCreationHistoryRunState string

const (
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStatefail          MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "Fail"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStatenotApplicable MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "NotApplicable"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStatepending       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "Pending"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStatescriptError   MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "ScriptError"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStatesuccess       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "Success"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunStateunknown       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "Unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppPolicyCreationHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStatefail),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStatenotApplicable),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStatepending),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStatescriptError),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStatesuccess),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunStateunknown),
	}
}

func parseMobileAppTroubleshootingAppPolicyCreationHistoryRunState(input string) (*MobileAppTroubleshootingAppPolicyCreationHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppPolicyCreationHistoryRunState{
		"fail":          MobileAppTroubleshootingAppPolicyCreationHistoryRunStatefail,
		"notapplicable": MobileAppTroubleshootingAppPolicyCreationHistoryRunStatenotApplicable,
		"pending":       MobileAppTroubleshootingAppPolicyCreationHistoryRunStatepending,
		"scripterror":   MobileAppTroubleshootingAppPolicyCreationHistoryRunStatescriptError,
		"success":       MobileAppTroubleshootingAppPolicyCreationHistoryRunStatesuccess,
		"unknown":       MobileAppTroubleshootingAppPolicyCreationHistoryRunStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppPolicyCreationHistoryRunState(input)
	return &out, nil
}
