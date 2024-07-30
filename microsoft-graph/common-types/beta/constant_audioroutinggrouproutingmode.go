package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioRoutingGroupRoutingMode string

const (
	AudioRoutingGroupRoutingMode_Multicast AudioRoutingGroupRoutingMode = "multicast"
	AudioRoutingGroupRoutingMode_OneToOne  AudioRoutingGroupRoutingMode = "oneToOne"
)

func PossibleValuesForAudioRoutingGroupRoutingMode() []string {
	return []string{
		string(AudioRoutingGroupRoutingMode_Multicast),
		string(AudioRoutingGroupRoutingMode_OneToOne),
	}
}

func (s *AudioRoutingGroupRoutingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAudioRoutingGroupRoutingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAudioRoutingGroupRoutingMode(input string) (*AudioRoutingGroupRoutingMode, error) {
	vals := map[string]AudioRoutingGroupRoutingMode{
		"multicast": AudioRoutingGroupRoutingMode_Multicast,
		"onetoone":  AudioRoutingGroupRoutingMode_OneToOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AudioRoutingGroupRoutingMode(input)
	return &out, nil
}
