package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppPolicyCreationHistoryRunState string

const (
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Fail          MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "fail"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_NotApplicable MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "notApplicable"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Pending       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "pending"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_ScriptError   MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "scriptError"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Success       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "success"
	MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Unknown       MobileAppTroubleshootingAppPolicyCreationHistoryRunState = "unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppPolicyCreationHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Fail),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_NotApplicable),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Pending),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_ScriptError),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Success),
		string(MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Unknown),
	}
}

func (s *MobileAppTroubleshootingAppPolicyCreationHistoryRunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppTroubleshootingAppPolicyCreationHistoryRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppTroubleshootingAppPolicyCreationHistoryRunState(input string) (*MobileAppTroubleshootingAppPolicyCreationHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppPolicyCreationHistoryRunState{
		"fail":          MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Fail,
		"notapplicable": MobileAppTroubleshootingAppPolicyCreationHistoryRunState_NotApplicable,
		"pending":       MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Pending,
		"scripterror":   MobileAppTroubleshootingAppPolicyCreationHistoryRunState_ScriptError,
		"success":       MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Success,
		"unknown":       MobileAppTroubleshootingAppPolicyCreationHistoryRunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppPolicyCreationHistoryRunState(input)
	return &out, nil
}
