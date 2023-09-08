package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingInfoRecordingStatus string

const (
	RecordingInfoRecordingStatusfailed       RecordingInfoRecordingStatus = "Failed"
	RecordingInfoRecordingStatusnotRecording RecordingInfoRecordingStatus = "NotRecording"
	RecordingInfoRecordingStatusrecording    RecordingInfoRecordingStatus = "Recording"
	RecordingInfoRecordingStatusunknown      RecordingInfoRecordingStatus = "Unknown"
)

func PossibleValuesForRecordingInfoRecordingStatus() []string {
	return []string{
		string(RecordingInfoRecordingStatusfailed),
		string(RecordingInfoRecordingStatusnotRecording),
		string(RecordingInfoRecordingStatusrecording),
		string(RecordingInfoRecordingStatusunknown),
	}
}

func parseRecordingInfoRecordingStatus(input string) (*RecordingInfoRecordingStatus, error) {
	vals := map[string]RecordingInfoRecordingStatus{
		"failed":       RecordingInfoRecordingStatusfailed,
		"notrecording": RecordingInfoRecordingStatusnotRecording,
		"recording":    RecordingInfoRecordingStatusrecording,
		"unknown":      RecordingInfoRecordingStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordingInfoRecordingStatus(input)
	return &out, nil
}
