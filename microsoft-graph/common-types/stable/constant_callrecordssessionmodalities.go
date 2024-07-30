package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsSessionModalities string

const (
	CallRecordsSessionModalities_Audio                   CallRecordsSessionModalities = "audio"
	CallRecordsSessionModalities_Data                    CallRecordsSessionModalities = "data"
	CallRecordsSessionModalities_ScreenSharing           CallRecordsSessionModalities = "screenSharing"
	CallRecordsSessionModalities_Video                   CallRecordsSessionModalities = "video"
	CallRecordsSessionModalities_VideoBasedScreenSharing CallRecordsSessionModalities = "videoBasedScreenSharing"
)

func PossibleValuesForCallRecordsSessionModalities() []string {
	return []string{
		string(CallRecordsSessionModalities_Audio),
		string(CallRecordsSessionModalities_Data),
		string(CallRecordsSessionModalities_ScreenSharing),
		string(CallRecordsSessionModalities_Video),
		string(CallRecordsSessionModalities_VideoBasedScreenSharing),
	}
}

func (s *CallRecordsSessionModalities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsSessionModalities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsSessionModalities(input string) (*CallRecordsSessionModalities, error) {
	vals := map[string]CallRecordsSessionModalities{
		"audio":                   CallRecordsSessionModalities_Audio,
		"data":                    CallRecordsSessionModalities_Data,
		"screensharing":           CallRecordsSessionModalities_ScreenSharing,
		"video":                   CallRecordsSessionModalities_Video,
		"videobasedscreensharing": CallRecordsSessionModalities_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsSessionModalities(input)
	return &out, nil
}
