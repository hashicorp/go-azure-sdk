package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppProductCodeDetectionProductVersionOperator string

const (
	Win32LobAppProductCodeDetectionProductVersionOperatorequal              Win32LobAppProductCodeDetectionProductVersionOperator = "Equal"
	Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThan        Win32LobAppProductCodeDetectionProductVersionOperator = "GreaterThan"
	Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThanOrEqual Win32LobAppProductCodeDetectionProductVersionOperator = "GreaterThanOrEqual"
	Win32LobAppProductCodeDetectionProductVersionOperatorlessThan           Win32LobAppProductCodeDetectionProductVersionOperator = "LessThan"
	Win32LobAppProductCodeDetectionProductVersionOperatorlessThanOrEqual    Win32LobAppProductCodeDetectionProductVersionOperator = "LessThanOrEqual"
	Win32LobAppProductCodeDetectionProductVersionOperatornotConfigured      Win32LobAppProductCodeDetectionProductVersionOperator = "NotConfigured"
	Win32LobAppProductCodeDetectionProductVersionOperatornotEqual           Win32LobAppProductCodeDetectionProductVersionOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppProductCodeDetectionProductVersionOperator() []string {
	return []string{
		string(Win32LobAppProductCodeDetectionProductVersionOperatorequal),
		string(Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThan),
		string(Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThanOrEqual),
		string(Win32LobAppProductCodeDetectionProductVersionOperatorlessThan),
		string(Win32LobAppProductCodeDetectionProductVersionOperatorlessThanOrEqual),
		string(Win32LobAppProductCodeDetectionProductVersionOperatornotConfigured),
		string(Win32LobAppProductCodeDetectionProductVersionOperatornotEqual),
	}
}

func parseWin32LobAppProductCodeDetectionProductVersionOperator(input string) (*Win32LobAppProductCodeDetectionProductVersionOperator, error) {
	vals := map[string]Win32LobAppProductCodeDetectionProductVersionOperator{
		"equal":              Win32LobAppProductCodeDetectionProductVersionOperatorequal,
		"greaterthan":        Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppProductCodeDetectionProductVersionOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppProductCodeDetectionProductVersionOperatorlessThan,
		"lessthanorequal":    Win32LobAppProductCodeDetectionProductVersionOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppProductCodeDetectionProductVersionOperatornotConfigured,
		"notequal":           Win32LobAppProductCodeDetectionProductVersionOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppProductCodeDetectionProductVersionOperator(input)
	return &out, nil
}
