package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryDetectionOperator string

const (
	Win32LobAppRegistryDetectionOperator_Equal              Win32LobAppRegistryDetectionOperator = "equal"
	Win32LobAppRegistryDetectionOperator_GreaterThan        Win32LobAppRegistryDetectionOperator = "greaterThan"
	Win32LobAppRegistryDetectionOperator_GreaterThanOrEqual Win32LobAppRegistryDetectionOperator = "greaterThanOrEqual"
	Win32LobAppRegistryDetectionOperator_LessThan           Win32LobAppRegistryDetectionOperator = "lessThan"
	Win32LobAppRegistryDetectionOperator_LessThanOrEqual    Win32LobAppRegistryDetectionOperator = "lessThanOrEqual"
	Win32LobAppRegistryDetectionOperator_NotConfigured      Win32LobAppRegistryDetectionOperator = "notConfigured"
	Win32LobAppRegistryDetectionOperator_NotEqual           Win32LobAppRegistryDetectionOperator = "notEqual"
)

func PossibleValuesForWin32LobAppRegistryDetectionOperator() []string {
	return []string{
		string(Win32LobAppRegistryDetectionOperator_Equal),
		string(Win32LobAppRegistryDetectionOperator_GreaterThan),
		string(Win32LobAppRegistryDetectionOperator_GreaterThanOrEqual),
		string(Win32LobAppRegistryDetectionOperator_LessThan),
		string(Win32LobAppRegistryDetectionOperator_LessThanOrEqual),
		string(Win32LobAppRegistryDetectionOperator_NotConfigured),
		string(Win32LobAppRegistryDetectionOperator_NotEqual),
	}
}

func (s *Win32LobAppRegistryDetectionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryDetectionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryDetectionOperator(input string) (*Win32LobAppRegistryDetectionOperator, error) {
	vals := map[string]Win32LobAppRegistryDetectionOperator{
		"equal":              Win32LobAppRegistryDetectionOperator_Equal,
		"greaterthan":        Win32LobAppRegistryDetectionOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppRegistryDetectionOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppRegistryDetectionOperator_LessThan,
		"lessthanorequal":    Win32LobAppRegistryDetectionOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppRegistryDetectionOperator_NotConfigured,
		"notequal":           Win32LobAppRegistryDetectionOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryDetectionOperator(input)
	return &out, nil
}
