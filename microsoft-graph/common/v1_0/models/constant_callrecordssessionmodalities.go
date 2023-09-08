package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsSessionModalities string

const (
	CallRecordsSessionModalitiesaudio                   CallRecordsSessionModalities = "Audio"
	CallRecordsSessionModalitiesdata                    CallRecordsSessionModalities = "Data"
	CallRecordsSessionModalitiesscreenSharing           CallRecordsSessionModalities = "ScreenSharing"
	CallRecordsSessionModalitiesvideo                   CallRecordsSessionModalities = "Video"
	CallRecordsSessionModalitiesvideoBasedScreenSharing CallRecordsSessionModalities = "VideoBasedScreenSharing"
)

func PossibleValuesForCallRecordsSessionModalities() []string {
	return []string{
		string(CallRecordsSessionModalitiesaudio),
		string(CallRecordsSessionModalitiesdata),
		string(CallRecordsSessionModalitiesscreenSharing),
		string(CallRecordsSessionModalitiesvideo),
		string(CallRecordsSessionModalitiesvideoBasedScreenSharing),
	}
}

func parseCallRecordsSessionModalities(input string) (*CallRecordsSessionModalities, error) {
	vals := map[string]CallRecordsSessionModalities{
		"audio":                   CallRecordsSessionModalitiesaudio,
		"data":                    CallRecordsSessionModalitiesdata,
		"screensharing":           CallRecordsSessionModalitiesscreenSharing,
		"video":                   CallRecordsSessionModalitiesvideo,
		"videobasedscreensharing": CallRecordsSessionModalitiesvideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsSessionModalities(input)
	return &out, nil
}
