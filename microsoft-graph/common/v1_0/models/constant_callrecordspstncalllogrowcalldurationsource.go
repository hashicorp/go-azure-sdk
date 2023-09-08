package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallLogRowCallDurationSource string

const (
	CallRecordsPstnCallLogRowCallDurationSourcemicrosoft CallRecordsPstnCallLogRowCallDurationSource = "Microsoft"
	CallRecordsPstnCallLogRowCallDurationSourceoperator  CallRecordsPstnCallLogRowCallDurationSource = "Operator"
)

func PossibleValuesForCallRecordsPstnCallLogRowCallDurationSource() []string {
	return []string{
		string(CallRecordsPstnCallLogRowCallDurationSourcemicrosoft),
		string(CallRecordsPstnCallLogRowCallDurationSourceoperator),
	}
}

func parseCallRecordsPstnCallLogRowCallDurationSource(input string) (*CallRecordsPstnCallLogRowCallDurationSource, error) {
	vals := map[string]CallRecordsPstnCallLogRowCallDurationSource{
		"microsoft": CallRecordsPstnCallLogRowCallDurationSourcemicrosoft,
		"operator":  CallRecordsPstnCallLogRowCallDurationSourceoperator,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnCallLogRowCallDurationSource(input)
	return &out, nil
}
