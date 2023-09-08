package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallState string

const (
	CallStateestablished      CallState = "Established"
	CallStateestablishing     CallState = "Establishing"
	CallStatehold             CallState = "Hold"
	CallStateincoming         CallState = "Incoming"
	CallStateredirecting      CallState = "Redirecting"
	CallStateringing          CallState = "Ringing"
	CallStateterminated       CallState = "Terminated"
	CallStateterminating      CallState = "Terminating"
	CallStatetransferAccepted CallState = "TransferAccepted"
	CallStatetransferring     CallState = "Transferring"
)

func PossibleValuesForCallState() []string {
	return []string{
		string(CallStateestablished),
		string(CallStateestablishing),
		string(CallStatehold),
		string(CallStateincoming),
		string(CallStateredirecting),
		string(CallStateringing),
		string(CallStateterminated),
		string(CallStateterminating),
		string(CallStatetransferAccepted),
		string(CallStatetransferring),
	}
}

func parseCallState(input string) (*CallState, error) {
	vals := map[string]CallState{
		"established":      CallStateestablished,
		"establishing":     CallStateestablishing,
		"hold":             CallStatehold,
		"incoming":         CallStateincoming,
		"redirecting":      CallStateredirecting,
		"ringing":          CallStateringing,
		"terminated":       CallStateterminated,
		"terminating":      CallStateterminating,
		"transferaccepted": CallStatetransferAccepted,
		"transferring":     CallStatetransferring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallState(input)
	return &out, nil
}
