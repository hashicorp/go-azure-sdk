package workspaceazureadonlyauthentications

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StateValue string

const (
	StateValueConsistent   StateValue = "Consistent"
	StateValueInConsistent StateValue = "InConsistent"
	StateValueUpdating     StateValue = "Updating"
)

func PossibleValuesForStateValue() []string {
	return []string{
		string(StateValueConsistent),
		string(StateValueInConsistent),
		string(StateValueUpdating),
	}
}

func parseStateValue(input string) (*StateValue, error) {
	vals := map[string]StateValue{
		"consistent":   StateValueConsistent,
		"inconsistent": StateValueInConsistent,
		"updating":     StateValueUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StateValue(input)
	return &out, nil
}
