package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppStateHistoryActionType string

const (
	MobileAppTroubleshootingAppStateHistoryActionTypeinstallCommandSent   MobileAppTroubleshootingAppStateHistoryActionType = "InstallCommandSent"
	MobileAppTroubleshootingAppStateHistoryActionTypeinstalled            MobileAppTroubleshootingAppStateHistoryActionType = "Installed"
	MobileAppTroubleshootingAppStateHistoryActionTypeuninstalled          MobileAppTroubleshootingAppStateHistoryActionType = "Uninstalled"
	MobileAppTroubleshootingAppStateHistoryActionTypeunknown              MobileAppTroubleshootingAppStateHistoryActionType = "Unknown"
	MobileAppTroubleshootingAppStateHistoryActionTypeuserRequestedInstall MobileAppTroubleshootingAppStateHistoryActionType = "UserRequestedInstall"
)

func PossibleValuesForMobileAppTroubleshootingAppStateHistoryActionType() []string {
	return []string{
		string(MobileAppTroubleshootingAppStateHistoryActionTypeinstallCommandSent),
		string(MobileAppTroubleshootingAppStateHistoryActionTypeinstalled),
		string(MobileAppTroubleshootingAppStateHistoryActionTypeuninstalled),
		string(MobileAppTroubleshootingAppStateHistoryActionTypeunknown),
		string(MobileAppTroubleshootingAppStateHistoryActionTypeuserRequestedInstall),
	}
}

func parseMobileAppTroubleshootingAppStateHistoryActionType(input string) (*MobileAppTroubleshootingAppStateHistoryActionType, error) {
	vals := map[string]MobileAppTroubleshootingAppStateHistoryActionType{
		"installcommandsent":   MobileAppTroubleshootingAppStateHistoryActionTypeinstallCommandSent,
		"installed":            MobileAppTroubleshootingAppStateHistoryActionTypeinstalled,
		"uninstalled":          MobileAppTroubleshootingAppStateHistoryActionTypeuninstalled,
		"unknown":              MobileAppTroubleshootingAppStateHistoryActionTypeunknown,
		"userrequestedinstall": MobileAppTroubleshootingAppStateHistoryActionTypeuserRequestedInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppTroubleshootingAppStateHistoryActionType(input)
	return &out, nil
}
