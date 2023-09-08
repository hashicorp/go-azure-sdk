package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRuleRuleType string

const (
	Win32LobAppRegistryRuleRuleTypedetection   Win32LobAppRegistryRuleRuleType = "Detection"
	Win32LobAppRegistryRuleRuleTyperequirement Win32LobAppRegistryRuleRuleType = "Requirement"
)

func PossibleValuesForWin32LobAppRegistryRuleRuleType() []string {
	return []string{
		string(Win32LobAppRegistryRuleRuleTypedetection),
		string(Win32LobAppRegistryRuleRuleTyperequirement),
	}
}

func parseWin32LobAppRegistryRuleRuleType(input string) (*Win32LobAppRegistryRuleRuleType, error) {
	vals := map[string]Win32LobAppRegistryRuleRuleType{
		"detection":   Win32LobAppRegistryRuleRuleTypedetection,
		"requirement": Win32LobAppRegistryRuleRuleTyperequirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRuleRuleType(input)
	return &out, nil
}
