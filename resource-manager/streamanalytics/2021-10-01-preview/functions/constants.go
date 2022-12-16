package functions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UdfType string

const (
	UdfTypeScalar UdfType = "Scalar"
)

func PossibleValuesForUdfType() []string {
	return []string{
		string(UdfTypeScalar),
	}
}

func parseUdfType(input string) (*UdfType, error) {
	vals := map[string]UdfType{
		"scalar": UdfTypeScalar,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UdfType(input)
	return &out, nil
}

type UpdateMode string

const (
	UpdateModeRefreshable UpdateMode = "Refreshable"
	UpdateModeStatic      UpdateMode = "Static"
)

func PossibleValuesForUpdateMode() []string {
	return []string{
		string(UpdateModeRefreshable),
		string(UpdateModeStatic),
	}
}

func parseUpdateMode(input string) (*UpdateMode, error) {
	vals := map[string]UpdateMode{
		"refreshable": UpdateModeRefreshable,
		"static":      UpdateModeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateMode(input)
	return &out, nil
}
