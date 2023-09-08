package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRequestedModalities string

const (
	CallRequestedModalitiesaudio                   CallRequestedModalities = "Audio"
	CallRequestedModalitiesdata                    CallRequestedModalities = "Data"
	CallRequestedModalitiesunknown                 CallRequestedModalities = "Unknown"
	CallRequestedModalitiesvideo                   CallRequestedModalities = "Video"
	CallRequestedModalitiesvideoBasedScreenSharing CallRequestedModalities = "VideoBasedScreenSharing"
)

func PossibleValuesForCallRequestedModalities() []string {
	return []string{
		string(CallRequestedModalitiesaudio),
		string(CallRequestedModalitiesdata),
		string(CallRequestedModalitiesunknown),
		string(CallRequestedModalitiesvideo),
		string(CallRequestedModalitiesvideoBasedScreenSharing),
	}
}

func parseCallRequestedModalities(input string) (*CallRequestedModalities, error) {
	vals := map[string]CallRequestedModalities{
		"audio":                   CallRequestedModalitiesaudio,
		"data":                    CallRequestedModalitiesdata,
		"unknown":                 CallRequestedModalitiesunknown,
		"video":                   CallRequestedModalitiesvideo,
		"videobasedscreensharing": CallRequestedModalitiesvideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRequestedModalities(input)
	return &out, nil
}
