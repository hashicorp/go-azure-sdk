package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallActiveModalities string

const (
	CallActiveModalities_Audio                   CallActiveModalities = "audio"
	CallActiveModalities_Data                    CallActiveModalities = "data"
	CallActiveModalities_Unknown                 CallActiveModalities = "unknown"
	CallActiveModalities_Video                   CallActiveModalities = "video"
	CallActiveModalities_VideoBasedScreenSharing CallActiveModalities = "videoBasedScreenSharing"
)

func PossibleValuesForCallActiveModalities() []string {
	return []string{
		string(CallActiveModalities_Audio),
		string(CallActiveModalities_Data),
		string(CallActiveModalities_Unknown),
		string(CallActiveModalities_Video),
		string(CallActiveModalities_VideoBasedScreenSharing),
	}
}

func (s *CallActiveModalities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallActiveModalities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallActiveModalities(input string) (*CallActiveModalities, error) {
	vals := map[string]CallActiveModalities{
		"audio":                   CallActiveModalities_Audio,
		"data":                    CallActiveModalities_Data,
		"unknown":                 CallActiveModalities_Unknown,
		"video":                   CallActiveModalities_Video,
		"videobasedscreensharing": CallActiveModalities_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallActiveModalities(input)
	return &out, nil
}
