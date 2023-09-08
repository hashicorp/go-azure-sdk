package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallRecordModalities string

const (
	CallRecordsCallRecordModalitiesaudio                   CallRecordsCallRecordModalities = "Audio"
	CallRecordsCallRecordModalitiesdata                    CallRecordsCallRecordModalities = "Data"
	CallRecordsCallRecordModalitiesscreenSharing           CallRecordsCallRecordModalities = "ScreenSharing"
	CallRecordsCallRecordModalitiesvideo                   CallRecordsCallRecordModalities = "Video"
	CallRecordsCallRecordModalitiesvideoBasedScreenSharing CallRecordsCallRecordModalities = "VideoBasedScreenSharing"
)

func PossibleValuesForCallRecordsCallRecordModalities() []string {
	return []string{
		string(CallRecordsCallRecordModalitiesaudio),
		string(CallRecordsCallRecordModalitiesdata),
		string(CallRecordsCallRecordModalitiesscreenSharing),
		string(CallRecordsCallRecordModalitiesvideo),
		string(CallRecordsCallRecordModalitiesvideoBasedScreenSharing),
	}
}

func parseCallRecordsCallRecordModalities(input string) (*CallRecordsCallRecordModalities, error) {
	vals := map[string]CallRecordsCallRecordModalities{
		"audio":                   CallRecordsCallRecordModalitiesaudio,
		"data":                    CallRecordsCallRecordModalitiesdata,
		"screensharing":           CallRecordsCallRecordModalitiesscreenSharing,
		"video":                   CallRecordsCallRecordModalitiesvideo,
		"videobasedscreensharing": CallRecordsCallRecordModalitiesvideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsCallRecordModalities(input)
	return &out, nil
}
