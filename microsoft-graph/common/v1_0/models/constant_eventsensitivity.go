package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSensitivity string

const (
	EventSensitivityconfidential EventSensitivity = "Confidential"
	EventSensitivitynormal       EventSensitivity = "Normal"
	EventSensitivitypersonal     EventSensitivity = "Personal"
	EventSensitivityprivate      EventSensitivity = "Private"
)

func PossibleValuesForEventSensitivity() []string {
	return []string{
		string(EventSensitivityconfidential),
		string(EventSensitivitynormal),
		string(EventSensitivitypersonal),
		string(EventSensitivityprivate),
	}
}

func parseEventSensitivity(input string) (*EventSensitivity, error) {
	vals := map[string]EventSensitivity{
		"confidential": EventSensitivityconfidential,
		"normal":       EventSensitivitynormal,
		"personal":     EventSensitivitypersonal,
		"private":      EventSensitivityprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSensitivity(input)
	return &out, nil
}
