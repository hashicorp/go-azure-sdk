package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeRuleProductVersionOperator string

const (
	Win32LobAppProductCodeRuleProductVersionOperator_Equal              Win32LobAppProductCodeRuleProductVersionOperator = "equal"
	Win32LobAppProductCodeRuleProductVersionOperator_GreaterThan        Win32LobAppProductCodeRuleProductVersionOperator = "greaterThan"
	Win32LobAppProductCodeRuleProductVersionOperator_GreaterThanOrEqual Win32LobAppProductCodeRuleProductVersionOperator = "greaterThanOrEqual"
	Win32LobAppProductCodeRuleProductVersionOperator_LessThan           Win32LobAppProductCodeRuleProductVersionOperator = "lessThan"
	Win32LobAppProductCodeRuleProductVersionOperator_LessThanOrEqual    Win32LobAppProductCodeRuleProductVersionOperator = "lessThanOrEqual"
	Win32LobAppProductCodeRuleProductVersionOperator_NotConfigured      Win32LobAppProductCodeRuleProductVersionOperator = "notConfigured"
	Win32LobAppProductCodeRuleProductVersionOperator_NotEqual           Win32LobAppProductCodeRuleProductVersionOperator = "notEqual"
)

func PossibleValuesForWin32LobAppProductCodeRuleProductVersionOperator() []string {
	return []string{
		string(Win32LobAppProductCodeRuleProductVersionOperator_Equal),
		string(Win32LobAppProductCodeRuleProductVersionOperator_GreaterThan),
		string(Win32LobAppProductCodeRuleProductVersionOperator_GreaterThanOrEqual),
		string(Win32LobAppProductCodeRuleProductVersionOperator_LessThan),
		string(Win32LobAppProductCodeRuleProductVersionOperator_LessThanOrEqual),
		string(Win32LobAppProductCodeRuleProductVersionOperator_NotConfigured),
		string(Win32LobAppProductCodeRuleProductVersionOperator_NotEqual),
	}
}

func (s *Win32LobAppProductCodeRuleProductVersionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppProductCodeRuleProductVersionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppProductCodeRuleProductVersionOperator(input string) (*Win32LobAppProductCodeRuleProductVersionOperator, error) {
	vals := map[string]Win32LobAppProductCodeRuleProductVersionOperator{
		"equal":              Win32LobAppProductCodeRuleProductVersionOperator_Equal,
		"greaterthan":        Win32LobAppProductCodeRuleProductVersionOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppProductCodeRuleProductVersionOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppProductCodeRuleProductVersionOperator_LessThan,
		"lessthanorequal":    Win32LobAppProductCodeRuleProductVersionOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppProductCodeRuleProductVersionOperator_NotConfigured,
		"notequal":           Win32LobAppProductCodeRuleProductVersionOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppProductCodeRuleProductVersionOperator(input)
	return &out, nil
}
