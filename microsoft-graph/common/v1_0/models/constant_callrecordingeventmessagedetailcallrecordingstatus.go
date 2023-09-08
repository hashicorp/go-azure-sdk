package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordingEventMessageDetailCallRecordingStatus string

const (
	CallRecordingEventMessageDetailCallRecordingStatuschunkFinished CallRecordingEventMessageDetailCallRecordingStatus = "ChunkFinished"
	CallRecordingEventMessageDetailCallRecordingStatusfailure       CallRecordingEventMessageDetailCallRecordingStatus = "Failure"
	CallRecordingEventMessageDetailCallRecordingStatusinitial       CallRecordingEventMessageDetailCallRecordingStatus = "Initial"
	CallRecordingEventMessageDetailCallRecordingStatussuccess       CallRecordingEventMessageDetailCallRecordingStatus = "Success"
)

func PossibleValuesForCallRecordingEventMessageDetailCallRecordingStatus() []string {
	return []string{
		string(CallRecordingEventMessageDetailCallRecordingStatuschunkFinished),
		string(CallRecordingEventMessageDetailCallRecordingStatusfailure),
		string(CallRecordingEventMessageDetailCallRecordingStatusinitial),
		string(CallRecordingEventMessageDetailCallRecordingStatussuccess),
	}
}

func parseCallRecordingEventMessageDetailCallRecordingStatus(input string) (*CallRecordingEventMessageDetailCallRecordingStatus, error) {
	vals := map[string]CallRecordingEventMessageDetailCallRecordingStatus{
		"chunkfinished": CallRecordingEventMessageDetailCallRecordingStatuschunkFinished,
		"failure":       CallRecordingEventMessageDetailCallRecordingStatusfailure,
		"initial":       CallRecordingEventMessageDetailCallRecordingStatusinitial,
		"success":       CallRecordingEventMessageDetailCallRecordingStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordingEventMessageDetailCallRecordingStatus(input)
	return &out, nil
}
