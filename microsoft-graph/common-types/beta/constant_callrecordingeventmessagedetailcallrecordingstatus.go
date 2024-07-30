package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordingEventMessageDetailCallRecordingStatus string

const (
	CallRecordingEventMessageDetailCallRecordingStatus_ChunkFinished CallRecordingEventMessageDetailCallRecordingStatus = "chunkFinished"
	CallRecordingEventMessageDetailCallRecordingStatus_Failure       CallRecordingEventMessageDetailCallRecordingStatus = "failure"
	CallRecordingEventMessageDetailCallRecordingStatus_Initial       CallRecordingEventMessageDetailCallRecordingStatus = "initial"
	CallRecordingEventMessageDetailCallRecordingStatus_Success       CallRecordingEventMessageDetailCallRecordingStatus = "success"
)

func PossibleValuesForCallRecordingEventMessageDetailCallRecordingStatus() []string {
	return []string{
		string(CallRecordingEventMessageDetailCallRecordingStatus_ChunkFinished),
		string(CallRecordingEventMessageDetailCallRecordingStatus_Failure),
		string(CallRecordingEventMessageDetailCallRecordingStatus_Initial),
		string(CallRecordingEventMessageDetailCallRecordingStatus_Success),
	}
}

func (s *CallRecordingEventMessageDetailCallRecordingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordingEventMessageDetailCallRecordingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordingEventMessageDetailCallRecordingStatus(input string) (*CallRecordingEventMessageDetailCallRecordingStatus, error) {
	vals := map[string]CallRecordingEventMessageDetailCallRecordingStatus{
		"chunkfinished": CallRecordingEventMessageDetailCallRecordingStatus_ChunkFinished,
		"failure":       CallRecordingEventMessageDetailCallRecordingStatus_Failure,
		"initial":       CallRecordingEventMessageDetailCallRecordingStatus_Initial,
		"success":       CallRecordingEventMessageDetailCallRecordingStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordingEventMessageDetailCallRecordingStatus(input)
	return &out, nil
}
