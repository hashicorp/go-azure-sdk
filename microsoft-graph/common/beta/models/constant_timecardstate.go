package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardState string

const (
	TimeCardStateclockedIn  TimeCardState = "ClockedIn"
	TimeCardStateclockedOut TimeCardState = "ClockedOut"
	TimeCardStateonBreak    TimeCardState = "OnBreak"
)

func PossibleValuesForTimeCardState() []string {
	return []string{
		string(TimeCardStateclockedIn),
		string(TimeCardStateclockedOut),
		string(TimeCardStateonBreak),
	}
}

func parseTimeCardState(input string) (*TimeCardState, error) {
	vals := map[string]TimeCardState{
		"clockedin":  TimeCardStateclockedIn,
		"clockedout": TimeCardStateclockedOut,
		"onbreak":    TimeCardStateonBreak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeCardState(input)
	return &out, nil
}
