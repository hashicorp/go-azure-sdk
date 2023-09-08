package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessControlItemAccessLevel string

const (
	WindowsPrivacyDataAccessControlItemAccessLevelforceAllow    WindowsPrivacyDataAccessControlItemAccessLevel = "ForceAllow"
	WindowsPrivacyDataAccessControlItemAccessLevelforceDeny     WindowsPrivacyDataAccessControlItemAccessLevel = "ForceDeny"
	WindowsPrivacyDataAccessControlItemAccessLevelnotConfigured WindowsPrivacyDataAccessControlItemAccessLevel = "NotConfigured"
	WindowsPrivacyDataAccessControlItemAccessLeveluserInControl WindowsPrivacyDataAccessControlItemAccessLevel = "UserInControl"
)

func PossibleValuesForWindowsPrivacyDataAccessControlItemAccessLevel() []string {
	return []string{
		string(WindowsPrivacyDataAccessControlItemAccessLevelforceAllow),
		string(WindowsPrivacyDataAccessControlItemAccessLevelforceDeny),
		string(WindowsPrivacyDataAccessControlItemAccessLevelnotConfigured),
		string(WindowsPrivacyDataAccessControlItemAccessLeveluserInControl),
	}
}

func parseWindowsPrivacyDataAccessControlItemAccessLevel(input string) (*WindowsPrivacyDataAccessControlItemAccessLevel, error) {
	vals := map[string]WindowsPrivacyDataAccessControlItemAccessLevel{
		"forceallow":    WindowsPrivacyDataAccessControlItemAccessLevelforceAllow,
		"forcedeny":     WindowsPrivacyDataAccessControlItemAccessLevelforceDeny,
		"notconfigured": WindowsPrivacyDataAccessControlItemAccessLevelnotConfigured,
		"userincontrol": WindowsPrivacyDataAccessControlItemAccessLeveluserInControl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessControlItemAccessLevel(input)
	return &out, nil
}
