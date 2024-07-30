package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingInfoRecordingStatus string

const (
	RecordingInfoRecordingStatus_Failed       RecordingInfoRecordingStatus = "failed"
	RecordingInfoRecordingStatus_NotRecording RecordingInfoRecordingStatus = "notRecording"
	RecordingInfoRecordingStatus_Recording    RecordingInfoRecordingStatus = "recording"
	RecordingInfoRecordingStatus_Unknown      RecordingInfoRecordingStatus = "unknown"
)

func PossibleValuesForRecordingInfoRecordingStatus() []string {
	return []string{
		string(RecordingInfoRecordingStatus_Failed),
		string(RecordingInfoRecordingStatus_NotRecording),
		string(RecordingInfoRecordingStatus_Recording),
		string(RecordingInfoRecordingStatus_Unknown),
	}
}

func (s *RecordingInfoRecordingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecordingInfoRecordingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecordingInfoRecordingStatus(input string) (*RecordingInfoRecordingStatus, error) {
	vals := map[string]RecordingInfoRecordingStatus{
		"failed":       RecordingInfoRecordingStatus_Failed,
		"notrecording": RecordingInfoRecordingStatus_NotRecording,
		"recording":    RecordingInfoRecordingStatus_Recording,
		"unknown":      RecordingInfoRecordingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordingInfoRecordingStatus(input)
	return &out, nil
}
