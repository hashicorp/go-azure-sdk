package sqlpoolsrestorepoints

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointType string

const (
	RestorePointTypeCONTINUOUS RestorePointType = "CONTINUOUS"
	RestorePointTypeDISCRETE   RestorePointType = "DISCRETE"
)

func PossibleValuesForRestorePointType() []string {
	return []string{
		string(RestorePointTypeCONTINUOUS),
		string(RestorePointTypeDISCRETE),
	}
}

func parseRestorePointType(input string) (*RestorePointType, error) {
	vals := map[string]RestorePointType{
		"continuous": RestorePointTypeCONTINUOUS,
		"discrete":   RestorePointTypeDISCRETE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestorePointType(input)
	return &out, nil
}
