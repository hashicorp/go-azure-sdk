package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallDirection string

const (
	CallDirectionincoming CallDirection = "Incoming"
	CallDirectionoutgoing CallDirection = "Outgoing"
)

func PossibleValuesForCallDirection() []string {
	return []string{
		string(CallDirectionincoming),
		string(CallDirectionoutgoing),
	}
}

func parseCallDirection(input string) (*CallDirection, error) {
	vals := map[string]CallDirection{
		"incoming": CallDirectionincoming,
		"outgoing": CallDirectionoutgoing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallDirection(input)
	return &out, nil
}
