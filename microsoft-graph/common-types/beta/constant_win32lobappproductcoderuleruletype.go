package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeRuleRuleType string

const (
	Win32LobAppProductCodeRuleRuleType_Detection   Win32LobAppProductCodeRuleRuleType = "detection"
	Win32LobAppProductCodeRuleRuleType_Requirement Win32LobAppProductCodeRuleRuleType = "requirement"
)

func PossibleValuesForWin32LobAppProductCodeRuleRuleType() []string {
	return []string{
		string(Win32LobAppProductCodeRuleRuleType_Detection),
		string(Win32LobAppProductCodeRuleRuleType_Requirement),
	}
}

func (s *Win32LobAppProductCodeRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppProductCodeRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppProductCodeRuleRuleType(input string) (*Win32LobAppProductCodeRuleRuleType, error) {
	vals := map[string]Win32LobAppProductCodeRuleRuleType{
		"detection":   Win32LobAppProductCodeRuleRuleType_Detection,
		"requirement": Win32LobAppProductCodeRuleRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppProductCodeRuleRuleType(input)
	return &out, nil
}
