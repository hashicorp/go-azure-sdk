package workspaceconnections

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValueFormat string

const (
	ValueFormatJSON ValueFormat = "JSON"
)

func PossibleValuesForValueFormat() []string {
	return []string{
		string(ValueFormatJSON),
	}
}

func parseValueFormat(input string) (*ValueFormat, error) {
	vals := map[string]ValueFormat{
		"json": ValueFormatJSON,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValueFormat(input)
	return &out, nil
}
