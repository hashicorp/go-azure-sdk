package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetectionOperator string

const (
	Win32LobAppFileSystemDetectionOperatorequal              Win32LobAppFileSystemDetectionOperator = "Equal"
	Win32LobAppFileSystemDetectionOperatorgreaterThan        Win32LobAppFileSystemDetectionOperator = "GreaterThan"
	Win32LobAppFileSystemDetectionOperatorgreaterThanOrEqual Win32LobAppFileSystemDetectionOperator = "GreaterThanOrEqual"
	Win32LobAppFileSystemDetectionOperatorlessThan           Win32LobAppFileSystemDetectionOperator = "LessThan"
	Win32LobAppFileSystemDetectionOperatorlessThanOrEqual    Win32LobAppFileSystemDetectionOperator = "LessThanOrEqual"
	Win32LobAppFileSystemDetectionOperatornotConfigured      Win32LobAppFileSystemDetectionOperator = "NotConfigured"
	Win32LobAppFileSystemDetectionOperatornotEqual           Win32LobAppFileSystemDetectionOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppFileSystemDetectionOperator() []string {
	return []string{
		string(Win32LobAppFileSystemDetectionOperatorequal),
		string(Win32LobAppFileSystemDetectionOperatorgreaterThan),
		string(Win32LobAppFileSystemDetectionOperatorgreaterThanOrEqual),
		string(Win32LobAppFileSystemDetectionOperatorlessThan),
		string(Win32LobAppFileSystemDetectionOperatorlessThanOrEqual),
		string(Win32LobAppFileSystemDetectionOperatornotConfigured),
		string(Win32LobAppFileSystemDetectionOperatornotEqual),
	}
}

func parseWin32LobAppFileSystemDetectionOperator(input string) (*Win32LobAppFileSystemDetectionOperator, error) {
	vals := map[string]Win32LobAppFileSystemDetectionOperator{
		"equal":              Win32LobAppFileSystemDetectionOperatorequal,
		"greaterthan":        Win32LobAppFileSystemDetectionOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppFileSystemDetectionOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppFileSystemDetectionOperatorlessThan,
		"lessthanorequal":    Win32LobAppFileSystemDetectionOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppFileSystemDetectionOperatornotConfigured,
		"notequal":           Win32LobAppFileSystemDetectionOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemDetectionOperator(input)
	return &out, nil
}
