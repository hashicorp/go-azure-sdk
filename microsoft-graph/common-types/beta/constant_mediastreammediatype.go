package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStreamMediaType string

const (
	MediaStreamMediaType_Audio                   MediaStreamMediaType = "audio"
	MediaStreamMediaType_Data                    MediaStreamMediaType = "data"
	MediaStreamMediaType_Unknown                 MediaStreamMediaType = "unknown"
	MediaStreamMediaType_Video                   MediaStreamMediaType = "video"
	MediaStreamMediaType_VideoBasedScreenSharing MediaStreamMediaType = "videoBasedScreenSharing"
)

func PossibleValuesForMediaStreamMediaType() []string {
	return []string{
		string(MediaStreamMediaType_Audio),
		string(MediaStreamMediaType_Data),
		string(MediaStreamMediaType_Unknown),
		string(MediaStreamMediaType_Video),
		string(MediaStreamMediaType_VideoBasedScreenSharing),
	}
}

func (s *MediaStreamMediaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaStreamMediaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaStreamMediaType(input string) (*MediaStreamMediaType, error) {
	vals := map[string]MediaStreamMediaType{
		"audio":                   MediaStreamMediaType_Audio,
		"data":                    MediaStreamMediaType_Data,
		"unknown":                 MediaStreamMediaType_Unknown,
		"video":                   MediaStreamMediaType_Video,
		"videobasedscreensharing": MediaStreamMediaType_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaStreamMediaType(input)
	return &out, nil
}
