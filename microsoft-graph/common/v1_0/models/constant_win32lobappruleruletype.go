package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRuleRuleType string

const (
	Win32LobAppRuleRuleTypedetection   Win32LobAppRuleRuleType = "Detection"
	Win32LobAppRuleRuleTyperequirement Win32LobAppRuleRuleType = "Requirement"
)

func PossibleValuesForWin32LobAppRuleRuleType() []string {
	return []string{
		string(Win32LobAppRuleRuleTypedetection),
		string(Win32LobAppRuleRuleTyperequirement),
	}
}

func parseWin32LobAppRuleRuleType(input string) (*Win32LobAppRuleRuleType, error) {
	vals := map[string]Win32LobAppRuleRuleType{
		"detection":   Win32LobAppRuleRuleTypedetection,
		"requirement": Win32LobAppRuleRuleTyperequirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRuleRuleType(input)
	return &out, nil
}
