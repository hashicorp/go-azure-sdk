package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallRecordModalities string

const (
	CallRecordsCallRecordModalities_Audio                   CallRecordsCallRecordModalities = "audio"
	CallRecordsCallRecordModalities_Data                    CallRecordsCallRecordModalities = "data"
	CallRecordsCallRecordModalities_ScreenSharing           CallRecordsCallRecordModalities = "screenSharing"
	CallRecordsCallRecordModalities_Video                   CallRecordsCallRecordModalities = "video"
	CallRecordsCallRecordModalities_VideoBasedScreenSharing CallRecordsCallRecordModalities = "videoBasedScreenSharing"
)

func PossibleValuesForCallRecordsCallRecordModalities() []string {
	return []string{
		string(CallRecordsCallRecordModalities_Audio),
		string(CallRecordsCallRecordModalities_Data),
		string(CallRecordsCallRecordModalities_ScreenSharing),
		string(CallRecordsCallRecordModalities_Video),
		string(CallRecordsCallRecordModalities_VideoBasedScreenSharing),
	}
}

func (s *CallRecordsCallRecordModalities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsCallRecordModalities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsCallRecordModalities(input string) (*CallRecordsCallRecordModalities, error) {
	vals := map[string]CallRecordsCallRecordModalities{
		"audio":                   CallRecordsCallRecordModalities_Audio,
		"data":                    CallRecordsCallRecordModalities_Data,
		"screensharing":           CallRecordsCallRecordModalities_ScreenSharing,
		"video":                   CallRecordsCallRecordModalities_Video,
		"videobasedscreensharing": CallRecordsCallRecordModalities_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsCallRecordModalities(input)
	return &out, nil
}
