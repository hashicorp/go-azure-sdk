package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeDetectionProductVersionOperator string

const (
	Win32LobAppProductCodeDetectionProductVersionOperator_Equal              Win32LobAppProductCodeDetectionProductVersionOperator = "equal"
	Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThan        Win32LobAppProductCodeDetectionProductVersionOperator = "greaterThan"
	Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThanOrEqual Win32LobAppProductCodeDetectionProductVersionOperator = "greaterThanOrEqual"
	Win32LobAppProductCodeDetectionProductVersionOperator_LessThan           Win32LobAppProductCodeDetectionProductVersionOperator = "lessThan"
	Win32LobAppProductCodeDetectionProductVersionOperator_LessThanOrEqual    Win32LobAppProductCodeDetectionProductVersionOperator = "lessThanOrEqual"
	Win32LobAppProductCodeDetectionProductVersionOperator_NotConfigured      Win32LobAppProductCodeDetectionProductVersionOperator = "notConfigured"
	Win32LobAppProductCodeDetectionProductVersionOperator_NotEqual           Win32LobAppProductCodeDetectionProductVersionOperator = "notEqual"
)

func PossibleValuesForWin32LobAppProductCodeDetectionProductVersionOperator() []string {
	return []string{
		string(Win32LobAppProductCodeDetectionProductVersionOperator_Equal),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThan),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThanOrEqual),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_LessThan),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_LessThanOrEqual),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_NotConfigured),
		string(Win32LobAppProductCodeDetectionProductVersionOperator_NotEqual),
	}
}

func (s *Win32LobAppProductCodeDetectionProductVersionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppProductCodeDetectionProductVersionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppProductCodeDetectionProductVersionOperator(input string) (*Win32LobAppProductCodeDetectionProductVersionOperator, error) {
	vals := map[string]Win32LobAppProductCodeDetectionProductVersionOperator{
		"equal":              Win32LobAppProductCodeDetectionProductVersionOperator_Equal,
		"greaterthan":        Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppProductCodeDetectionProductVersionOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppProductCodeDetectionProductVersionOperator_LessThan,
		"lessthanorequal":    Win32LobAppProductCodeDetectionProductVersionOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppProductCodeDetectionProductVersionOperator_NotConfigured,
		"notequal":           Win32LobAppProductCodeDetectionProductVersionOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppProductCodeDetectionProductVersionOperator(input)
	return &out, nil
}
