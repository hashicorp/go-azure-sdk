package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallActiveModalities string

const (
	CallActiveModalitiesaudio                   CallActiveModalities = "Audio"
	CallActiveModalitiesdata                    CallActiveModalities = "Data"
	CallActiveModalitiesunknown                 CallActiveModalities = "Unknown"
	CallActiveModalitiesvideo                   CallActiveModalities = "Video"
	CallActiveModalitiesvideoBasedScreenSharing CallActiveModalities = "VideoBasedScreenSharing"
)

func PossibleValuesForCallActiveModalities() []string {
	return []string{
		string(CallActiveModalitiesaudio),
		string(CallActiveModalitiesdata),
		string(CallActiveModalitiesunknown),
		string(CallActiveModalitiesvideo),
		string(CallActiveModalitiesvideoBasedScreenSharing),
	}
}

func parseCallActiveModalities(input string) (*CallActiveModalities, error) {
	vals := map[string]CallActiveModalities{
		"audio":                   CallActiveModalitiesaudio,
		"data":                    CallActiveModalitiesdata,
		"unknown":                 CallActiveModalitiesunknown,
		"video":                   CallActiveModalitiesvideo,
		"videobasedscreensharing": CallActiveModalitiesvideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallActiveModalities(input)
	return &out, nil
}
