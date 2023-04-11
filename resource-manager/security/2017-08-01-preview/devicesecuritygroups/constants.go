package devicesecuritygroups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValueType string

const (
	ValueTypeIPCidr ValueType = "IpCidr"
	ValueTypeString ValueType = "String"
)

func PossibleValuesForValueType() []string {
	return []string{
		string(ValueTypeIPCidr),
		string(ValueTypeString),
	}
}

func parseValueType(input string) (*ValueType, error) {
	vals := map[string]ValueType{
		"ipcidr": ValueTypeIPCidr,
		"string": ValueTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValueType(input)
	return &out, nil
}
