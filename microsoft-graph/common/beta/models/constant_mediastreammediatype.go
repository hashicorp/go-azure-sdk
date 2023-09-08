package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStreamMediaType string

const (
	MediaStreamMediaTypeaudio                   MediaStreamMediaType = "Audio"
	MediaStreamMediaTypedata                    MediaStreamMediaType = "Data"
	MediaStreamMediaTypeunknown                 MediaStreamMediaType = "Unknown"
	MediaStreamMediaTypevideo                   MediaStreamMediaType = "Video"
	MediaStreamMediaTypevideoBasedScreenSharing MediaStreamMediaType = "VideoBasedScreenSharing"
)

func PossibleValuesForMediaStreamMediaType() []string {
	return []string{
		string(MediaStreamMediaTypeaudio),
		string(MediaStreamMediaTypedata),
		string(MediaStreamMediaTypeunknown),
		string(MediaStreamMediaTypevideo),
		string(MediaStreamMediaTypevideoBasedScreenSharing),
	}
}

func parseMediaStreamMediaType(input string) (*MediaStreamMediaType, error) {
	vals := map[string]MediaStreamMediaType{
		"audio":                   MediaStreamMediaTypeaudio,
		"data":                    MediaStreamMediaTypedata,
		"unknown":                 MediaStreamMediaTypeunknown,
		"video":                   MediaStreamMediaTypevideo,
		"videobasedscreensharing": MediaStreamMediaTypevideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaStreamMediaType(input)
	return &out, nil
}
