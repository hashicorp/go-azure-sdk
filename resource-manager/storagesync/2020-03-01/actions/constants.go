package actions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDetectionMode string

const (
	ChangeDetectionModeDefault   ChangeDetectionMode = "Default"
	ChangeDetectionModeRecursive ChangeDetectionMode = "Recursive"
)

func PossibleValuesForChangeDetectionMode() []string {
	return []string{
		string(ChangeDetectionModeDefault),
		string(ChangeDetectionModeRecursive),
	}
}

func parseChangeDetectionMode(input string) (*ChangeDetectionMode, error) {
	vals := map[string]ChangeDetectionMode{
		"default":   ChangeDetectionModeDefault,
		"recursive": ChangeDetectionModeRecursive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDetectionMode(input)
	return &out, nil
}
