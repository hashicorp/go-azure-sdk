package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeRuleRuleType string

const (
	Win32LobAppProductCodeRuleRuleTypedetection   Win32LobAppProductCodeRuleRuleType = "Detection"
	Win32LobAppProductCodeRuleRuleTyperequirement Win32LobAppProductCodeRuleRuleType = "Requirement"
)

func PossibleValuesForWin32LobAppProductCodeRuleRuleType() []string {
	return []string{
		string(Win32LobAppProductCodeRuleRuleTypedetection),
		string(Win32LobAppProductCodeRuleRuleTyperequirement),
	}
}

func parseWin32LobAppProductCodeRuleRuleType(input string) (*Win32LobAppProductCodeRuleRuleType, error) {
	vals := map[string]Win32LobAppProductCodeRuleRuleType{
		"detection":   Win32LobAppProductCodeRuleRuleTypedetection,
		"requirement": Win32LobAppProductCodeRuleRuleTyperequirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppProductCodeRuleRuleType(input)
	return &out, nil
}
