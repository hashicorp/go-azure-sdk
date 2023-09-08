package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallMediaStateAudio string

const (
	CallMediaStateAudioactive   CallMediaStateAudio = "Active"
	CallMediaStateAudioinactive CallMediaStateAudio = "Inactive"
)

func PossibleValuesForCallMediaStateAudio() []string {
	return []string{
		string(CallMediaStateAudioactive),
		string(CallMediaStateAudioinactive),
	}
}

func parseCallMediaStateAudio(input string) (*CallMediaStateAudio, error) {
	vals := map[string]CallMediaStateAudio{
		"active":   CallMediaStateAudioactive,
		"inactive": CallMediaStateAudioinactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallMediaStateAudio(input)
	return &out, nil
}
