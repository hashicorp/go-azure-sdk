package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRuleRuleType string

const (
	Win32LobAppRegistryRuleRuleType_Detection   Win32LobAppRegistryRuleRuleType = "detection"
	Win32LobAppRegistryRuleRuleType_Requirement Win32LobAppRegistryRuleRuleType = "requirement"
)

func PossibleValuesForWin32LobAppRegistryRuleRuleType() []string {
	return []string{
		string(Win32LobAppRegistryRuleRuleType_Detection),
		string(Win32LobAppRegistryRuleRuleType_Requirement),
	}
}

func (s *Win32LobAppRegistryRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryRuleRuleType(input string) (*Win32LobAppRegistryRuleRuleType, error) {
	vals := map[string]Win32LobAppRegistryRuleRuleType{
		"detection":   Win32LobAppRegistryRuleRuleType_Detection,
		"requirement": Win32LobAppRegistryRuleRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRuleRuleType(input)
	return &out, nil
}
