package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppStateHistoryRunState string

const (
	MobileAppTroubleshootingAppStateHistoryRunState_Fail          MobileAppTroubleshootingAppStateHistoryRunState = "fail"
	MobileAppTroubleshootingAppStateHistoryRunState_NotApplicable MobileAppTroubleshootingAppStateHistoryRunState = "notApplicable"
	MobileAppTroubleshootingAppStateHistoryRunState_Pending       MobileAppTroubleshootingAppStateHistoryRunState = "pending"
	MobileAppTroubleshootingAppStateHistoryRunState_ScriptError   MobileAppTroubleshootingAppStateHistoryRunState = "scriptError"
	MobileAppTroubleshootingAppStateHistoryRunState_Success       MobileAppTroubleshootingAppStateHistoryRunState = "success"
	MobileAppTroubleshootingAppStateHistoryRunState_Unknown       MobileAppTroubleshootingAppStateHistoryRunState = "unknown"
)

func PossibleValuesForMobileAppTroubleshootingAppStateHistoryRunState() []string {
	return []string{
		string(MobileAppTroubleshootingAppStateHistoryRunState_Fail),
		string(MobileAppTroubleshootingAppStateHistoryRunState_NotApplicable),
		string(MobileAppTroubleshootingAppStateHistoryRunState_Pending),
		string(MobileAppTroubleshootingAppStateHistoryRunState_ScriptError),
		string(MobileAppTroubleshootingAppStateHistoryRunState_Success),
		string(MobileAppTroubleshootingAppStateHistoryRunState_Unknown),
	}
}

func (s *MobileAppTroubleshootingAppStateHistoryRunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppTroubleshootingAppStateHistoryRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppTroubleshootingAppStateHistoryRunState(input string) (*MobileAppTroubleshootingAppStateHistoryRunState, error) {
	vals := map[string]MobileAppTroubleshootingAppStateHistoryRunState{
		"fail":          MobileAppTroubleshootingAppStateHistoryRunState_Fail,
		"notapplicable": MobileAppTroubleshootingAppStateHistoryRunState_NotApplicable,
		"pending":       MobileAppTroubleshootingAppStateHistoryRunState_Pending,
		"scripterror":   MobileAppTroubleshootingAppStateHistoryRunState_ScriptError,
		"success":       MobileAppTroubleshootingAppStateHistoryRunState_Success,
		"unknown":       MobileAppTroubleshootingAppStateHistoryRunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppStateHistoryRunState(input)
	return &out, nil
}
