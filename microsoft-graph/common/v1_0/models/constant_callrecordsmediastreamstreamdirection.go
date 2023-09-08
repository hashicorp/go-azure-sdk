package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamStreamDirection string

const (
	CallRecordsMediaStreamStreamDirectioncalleeToCaller CallRecordsMediaStreamStreamDirection = "CalleeToCaller"
	CallRecordsMediaStreamStreamDirectioncallerToCallee CallRecordsMediaStreamStreamDirection = "CallerToCallee"
)

func PossibleValuesForCallRecordsMediaStreamStreamDirection() []string {
	return []string{
		string(CallRecordsMediaStreamStreamDirectioncalleeToCaller),
		string(CallRecordsMediaStreamStreamDirectioncallerToCallee),
	}
}

func parseCallRecordsMediaStreamStreamDirection(input string) (*CallRecordsMediaStreamStreamDirection, error) {
	vals := map[string]CallRecordsMediaStreamStreamDirection{
		"calleetocaller": CallRecordsMediaStreamStreamDirectioncalleeToCaller,
		"callertocallee": CallRecordsMediaStreamStreamDirectioncallerToCallee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamStreamDirection(input)
	return &out, nil
}
