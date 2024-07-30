package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRequestedModalities string

const (
	CallRequestedModalities_Audio                   CallRequestedModalities = "audio"
	CallRequestedModalities_Data                    CallRequestedModalities = "data"
	CallRequestedModalities_Video                   CallRequestedModalities = "video"
	CallRequestedModalities_VideoBasedScreenSharing CallRequestedModalities = "videoBasedScreenSharing"
)

func PossibleValuesForCallRequestedModalities() []string {
	return []string{
		string(CallRequestedModalities_Audio),
		string(CallRequestedModalities_Data),
		string(CallRequestedModalities_Video),
		string(CallRequestedModalities_VideoBasedScreenSharing),
	}
}

func (s *CallRequestedModalities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRequestedModalities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRequestedModalities(input string) (*CallRequestedModalities, error) {
	vals := map[string]CallRequestedModalities{
		"audio":                   CallRequestedModalities_Audio,
		"data":                    CallRequestedModalities_Data,
		"video":                   CallRequestedModalities_Video,
		"videobasedscreensharing": CallRequestedModalities_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRequestedModalities(input)
	return &out, nil
}
