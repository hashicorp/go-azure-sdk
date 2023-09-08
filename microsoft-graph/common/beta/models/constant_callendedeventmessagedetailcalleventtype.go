package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEndedEventMessageDetailCallEventType string

const (
	CallEndedEventMessageDetailCallEventTypecall        CallEndedEventMessageDetailCallEventType = "Call"
	CallEndedEventMessageDetailCallEventTypemeeting     CallEndedEventMessageDetailCallEventType = "Meeting"
	CallEndedEventMessageDetailCallEventTypescreenShare CallEndedEventMessageDetailCallEventType = "ScreenShare"
)

func PossibleValuesForCallEndedEventMessageDetailCallEventType() []string {
	return []string{
		string(CallEndedEventMessageDetailCallEventTypecall),
		string(CallEndedEventMessageDetailCallEventTypemeeting),
		string(CallEndedEventMessageDetailCallEventTypescreenShare),
	}
}

func parseCallEndedEventMessageDetailCallEventType(input string) (*CallEndedEventMessageDetailCallEventType, error) {
	vals := map[string]CallEndedEventMessageDetailCallEventType{
		"call":        CallEndedEventMessageDetailCallEventTypecall,
		"meeting":     CallEndedEventMessageDetailCallEventTypemeeting,
		"screenshare": CallEndedEventMessageDetailCallEventTypescreenShare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallEndedEventMessageDetailCallEventType(input)
	return &out, nil
}
