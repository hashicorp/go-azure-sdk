package hypervmachines

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HighlyAvailable string

const (
	HighlyAvailableNo      HighlyAvailable = "No"
	HighlyAvailableUnknown HighlyAvailable = "Unknown"
	HighlyAvailableYes     HighlyAvailable = "Yes"
)

func PossibleValuesForHighlyAvailable() []string {
	return []string{
		string(HighlyAvailableNo),
		string(HighlyAvailableUnknown),
		string(HighlyAvailableYes),
	}
}

func parseHighlyAvailable(input string) (*HighlyAvailable, error) {
	vals := map[string]HighlyAvailable{
		"no":      HighlyAvailableNo,
		"unknown": HighlyAvailableUnknown,
		"yes":     HighlyAvailableYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HighlyAvailable(input)
	return &out, nil
}
