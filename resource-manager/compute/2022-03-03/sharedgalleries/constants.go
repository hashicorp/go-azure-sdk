package sharedgalleries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedToValues string

const (
	SharedToValuesTenant SharedToValues = "tenant"
)

func PossibleValuesForSharedToValues() []string {
	return []string{
		string(SharedToValuesTenant),
	}
}

func parseSharedToValues(input string) (*SharedToValues, error) {
	vals := map[string]SharedToValues{
		"tenant": SharedToValuesTenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedToValues(input)
	return &out, nil
}
