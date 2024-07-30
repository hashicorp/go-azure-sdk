package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetectionOperator string

const (
	Win32LobAppFileSystemDetectionOperator_Equal              Win32LobAppFileSystemDetectionOperator = "equal"
	Win32LobAppFileSystemDetectionOperator_GreaterThan        Win32LobAppFileSystemDetectionOperator = "greaterThan"
	Win32LobAppFileSystemDetectionOperator_GreaterThanOrEqual Win32LobAppFileSystemDetectionOperator = "greaterThanOrEqual"
	Win32LobAppFileSystemDetectionOperator_LessThan           Win32LobAppFileSystemDetectionOperator = "lessThan"
	Win32LobAppFileSystemDetectionOperator_LessThanOrEqual    Win32LobAppFileSystemDetectionOperator = "lessThanOrEqual"
	Win32LobAppFileSystemDetectionOperator_NotConfigured      Win32LobAppFileSystemDetectionOperator = "notConfigured"
	Win32LobAppFileSystemDetectionOperator_NotEqual           Win32LobAppFileSystemDetectionOperator = "notEqual"
)

func PossibleValuesForWin32LobAppFileSystemDetectionOperator() []string {
	return []string{
		string(Win32LobAppFileSystemDetectionOperator_Equal),
		string(Win32LobAppFileSystemDetectionOperator_GreaterThan),
		string(Win32LobAppFileSystemDetectionOperator_GreaterThanOrEqual),
		string(Win32LobAppFileSystemDetectionOperator_LessThan),
		string(Win32LobAppFileSystemDetectionOperator_LessThanOrEqual),
		string(Win32LobAppFileSystemDetectionOperator_NotConfigured),
		string(Win32LobAppFileSystemDetectionOperator_NotEqual),
	}
}

func (s *Win32LobAppFileSystemDetectionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemDetectionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemDetectionOperator(input string) (*Win32LobAppFileSystemDetectionOperator, error) {
	vals := map[string]Win32LobAppFileSystemDetectionOperator{
		"equal":              Win32LobAppFileSystemDetectionOperator_Equal,
		"greaterthan":        Win32LobAppFileSystemDetectionOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppFileSystemDetectionOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppFileSystemDetectionOperator_LessThan,
		"lessthanorequal":    Win32LobAppFileSystemDetectionOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppFileSystemDetectionOperator_NotConfigured,
		"notequal":           Win32LobAppFileSystemDetectionOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemDetectionOperator(input)
	return &out, nil
}
