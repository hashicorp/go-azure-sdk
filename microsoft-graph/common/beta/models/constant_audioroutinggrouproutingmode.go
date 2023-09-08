package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioRoutingGroupRoutingMode string

const (
	AudioRoutingGroupRoutingModemulticast AudioRoutingGroupRoutingMode = "Multicast"
	AudioRoutingGroupRoutingModeoneToOne  AudioRoutingGroupRoutingMode = "OneToOne"
)

func PossibleValuesForAudioRoutingGroupRoutingMode() []string {
	return []string{
		string(AudioRoutingGroupRoutingModemulticast),
		string(AudioRoutingGroupRoutingModeoneToOne),
	}
}

func parseAudioRoutingGroupRoutingMode(input string) (*AudioRoutingGroupRoutingMode, error) {
	vals := map[string]AudioRoutingGroupRoutingMode{
		"multicast": AudioRoutingGroupRoutingModemulticast,
		"onetoone":  AudioRoutingGroupRoutingModeoneToOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AudioRoutingGroupRoutingMode(input)
	return &out, nil
}
