package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRuleRuleType string

const (
	Win32LobAppFileSystemRuleRuleType_Detection   Win32LobAppFileSystemRuleRuleType = "detection"
	Win32LobAppFileSystemRuleRuleType_Requirement Win32LobAppFileSystemRuleRuleType = "requirement"
)

func PossibleValuesForWin32LobAppFileSystemRuleRuleType() []string {
	return []string{
		string(Win32LobAppFileSystemRuleRuleType_Detection),
		string(Win32LobAppFileSystemRuleRuleType_Requirement),
	}
}

func (s *Win32LobAppFileSystemRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemRuleRuleType(input string) (*Win32LobAppFileSystemRuleRuleType, error) {
	vals := map[string]Win32LobAppFileSystemRuleRuleType{
		"detection":   Win32LobAppFileSystemRuleRuleType_Detection,
		"requirement": Win32LobAppFileSystemRuleRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRuleRuleType(input)
	return &out, nil
}
