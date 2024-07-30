package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRequirementOperator string

const (
	Win32LobAppFileSystemRequirementOperator_Equal              Win32LobAppFileSystemRequirementOperator = "equal"
	Win32LobAppFileSystemRequirementOperator_GreaterThan        Win32LobAppFileSystemRequirementOperator = "greaterThan"
	Win32LobAppFileSystemRequirementOperator_GreaterThanOrEqual Win32LobAppFileSystemRequirementOperator = "greaterThanOrEqual"
	Win32LobAppFileSystemRequirementOperator_LessThan           Win32LobAppFileSystemRequirementOperator = "lessThan"
	Win32LobAppFileSystemRequirementOperator_LessThanOrEqual    Win32LobAppFileSystemRequirementOperator = "lessThanOrEqual"
	Win32LobAppFileSystemRequirementOperator_NotConfigured      Win32LobAppFileSystemRequirementOperator = "notConfigured"
	Win32LobAppFileSystemRequirementOperator_NotEqual           Win32LobAppFileSystemRequirementOperator = "notEqual"
)

func PossibleValuesForWin32LobAppFileSystemRequirementOperator() []string {
	return []string{
		string(Win32LobAppFileSystemRequirementOperator_Equal),
		string(Win32LobAppFileSystemRequirementOperator_GreaterThan),
		string(Win32LobAppFileSystemRequirementOperator_GreaterThanOrEqual),
		string(Win32LobAppFileSystemRequirementOperator_LessThan),
		string(Win32LobAppFileSystemRequirementOperator_LessThanOrEqual),
		string(Win32LobAppFileSystemRequirementOperator_NotConfigured),
		string(Win32LobAppFileSystemRequirementOperator_NotEqual),
	}
}

func (s *Win32LobAppFileSystemRequirementOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemRequirementOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemRequirementOperator(input string) (*Win32LobAppFileSystemRequirementOperator, error) {
	vals := map[string]Win32LobAppFileSystemRequirementOperator{
		"equal":              Win32LobAppFileSystemRequirementOperator_Equal,
		"greaterthan":        Win32LobAppFileSystemRequirementOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppFileSystemRequirementOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppFileSystemRequirementOperator_LessThan,
		"lessthanorequal":    Win32LobAppFileSystemRequirementOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppFileSystemRequirementOperator_NotConfigured,
		"notequal":           Win32LobAppFileSystemRequirementOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRequirementOperator(input)
	return &out, nil
}
