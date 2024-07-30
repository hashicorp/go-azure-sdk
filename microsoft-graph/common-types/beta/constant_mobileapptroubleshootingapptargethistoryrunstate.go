package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppTargetHistoryRunState string

const (
	MobileAppTroubleshootingAppTargetHistoryRunState_Fail          MobileAppTroubleshootingAppTargetHistoryRunState = "fail"
	MobileAppTroubleshootingAppTargetHistoryRunState_NotApplicable MobileAppTroubleshootingAppTargetHistoryRunState = "notApplicable"
	MobileAppTroubleshootingAppTargetHistoryRunState_Pending       MobileAppTroubleshootingAppTargetHistoryRunState = "pending"
	MobileAppTroubleshootingAppTargetHistoryRunState_ScriptError   MobileAppTroubleshootingAppTargetHistoryRunState = "scriptError"
	MobileAppTroubleshootingAppTargetHistoryRunState_Success       MobileAppTroubleshootingAppTargetHistoryRunState = "success"
	MobileAppTroubleshootingAppTargetHistoryRunState_Unknown       MobileAppTroubleshootingAppTargetHistoryRunState = "unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppTargetHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppTargetHistoryRunState_Fail),
		string(MobileAppTroubleshootingAppTargetHistoryRunState_NotApplicable),
		string(MobileAppTroubleshootingAppTargetHistoryRunState_Pending),
		string(MobileAppTroubleshootingAppTargetHistoryRunState_ScriptError),
		string(MobileAppTroubleshootingAppTargetHistoryRunState_Success),
		string(MobileAppTroubleshootingAppTargetHistoryRunState_Unknown),
	}
}

func (s *MobileAppTroubleshootingAppTargetHistoryRunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppTroubleshootingAppTargetHistoryRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppTroubleshootingAppTargetHistoryRunState(input string) (*MobileAppTroubleshootingAppTargetHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppTargetHistoryRunState{
		"fail":          MobileAppTroubleshootingAppTargetHistoryRunState_Fail,
		"notapplicable": MobileAppTroubleshootingAppTargetHistoryRunState_NotApplicable,
		"pending":       MobileAppTroubleshootingAppTargetHistoryRunState_Pending,
		"scripterror":   MobileAppTroubleshootingAppTargetHistoryRunState_ScriptError,
		"success":       MobileAppTroubleshootingAppTargetHistoryRunState_Success,
		"unknown":       MobileAppTroubleshootingAppTargetHistoryRunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppTargetHistoryRunState(input)
	return &out, nil
}
