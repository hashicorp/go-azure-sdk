package applicationpackage

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageState string

const (
	PackageStateActive  PackageState = "Active"
	PackageStatePending PackageState = "Pending"
)

func PossibleValuesForPackageState() []string {
	return []string{
		string(PackageStateActive),
		string(PackageStatePending),
	}
}

func parsePackageState(input string) (*PackageState, error) {
	vals := map[string]PackageState{
		"active":  PackageStateActive,
		"pending": PackageStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PackageState(input)
	return &out, nil
}
