package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRuleRuleType string

const (
	Win32LobAppRuleRuleType_Detection   Win32LobAppRuleRuleType = "detection"
	Win32LobAppRuleRuleType_Requirement Win32LobAppRuleRuleType = "requirement"
)

func PossibleValuesForWin32LobAppRuleRuleType() []string {
	return []string{
		string(Win32LobAppRuleRuleType_Detection),
		string(Win32LobAppRuleRuleType_Requirement),
	}
}

func (s *Win32LobAppRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRuleRuleType(input string) (*Win32LobAppRuleRuleType, error) {
	vals := map[string]Win32LobAppRuleRuleType{
		"detection":   Win32LobAppRuleRuleType_Detection,
		"requirement": Win32LobAppRuleRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRuleRuleType(input)
	return &out, nil
}
