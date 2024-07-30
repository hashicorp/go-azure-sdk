package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppStateHistoryActionType string

const (
	MobileAppTroubleshootingAppStateHistoryActionType_InstallCommandSent   MobileAppTroubleshootingAppStateHistoryActionType = "installCommandSent"
	MobileAppTroubleshootingAppStateHistoryActionType_Installed            MobileAppTroubleshootingAppStateHistoryActionType = "installed"
	MobileAppTroubleshootingAppStateHistoryActionType_Uninstalled          MobileAppTroubleshootingAppStateHistoryActionType = "uninstalled"
	MobileAppTroubleshootingAppStateHistoryActionType_Unknown              MobileAppTroubleshootingAppStateHistoryActionType = "unknown"
	MobileAppTroubleshootingAppStateHistoryActionType_UserRequestedInstall MobileAppTroubleshootingAppStateHistoryActionType = "userRequestedInstall"
)

func PossibleValuesForMobileAppTroubleshootingAppStateHistoryActionType() []string {
	return []string{
		string(MobileAppTroubleshootingAppStateHistoryActionType_InstallCommandSent),
		string(MobileAppTroubleshootingAppStateHistoryActionType_Installed),
		string(MobileAppTroubleshootingAppStateHistoryActionType_Uninstalled),
		string(MobileAppTroubleshootingAppStateHistoryActionType_Unknown),
		string(MobileAppTroubleshootingAppStateHistoryActionType_UserRequestedInstall),
	}
}

func (s *MobileAppTroubleshootingAppStateHistoryActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppTroubleshootingAppStateHistoryActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppTroubleshootingAppStateHistoryActionType(input string) (*MobileAppTroubleshootingAppStateHistoryActionType, error) {
	vals := map[string]MobileAppTroubleshootingAppStateHistoryActionType{
		"installcommandsent":   MobileAppTroubleshootingAppStateHistoryActionType_InstallCommandSent,
		"installed":            MobileAppTroubleshootingAppStateHistoryActionType_Installed,
		"uninstalled":          MobileAppTroubleshootingAppStateHistoryActionType_Uninstalled,
		"unknown":              MobileAppTroubleshootingAppStateHistoryActionType_Unknown,
		"userrequestedinstall": MobileAppTroubleshootingAppStateHistoryActionType_UserRequestedInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppStateHistoryActionType(input)
	return &out, nil
}
