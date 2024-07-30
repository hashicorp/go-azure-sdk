package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaSourceContentCategory string

const (
	MediaSourceContentCategory_LiveStream      MediaSourceContentCategory = "liveStream"
	MediaSourceContentCategory_Meeting         MediaSourceContentCategory = "meeting"
	MediaSourceContentCategory_Presentation    MediaSourceContentCategory = "presentation"
	MediaSourceContentCategory_ScreenRecording MediaSourceContentCategory = "screenRecording"
)

func PossibleValuesForMediaSourceContentCategory() []string {
	return []string{
		string(MediaSourceContentCategory_LiveStream),
		string(MediaSourceContentCategory_Meeting),
		string(MediaSourceContentCategory_Presentation),
		string(MediaSourceContentCategory_ScreenRecording),
	}
}

func (s *MediaSourceContentCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaSourceContentCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaSourceContentCategory(input string) (*MediaSourceContentCategory, error) {
	vals := map[string]MediaSourceContentCategory{
		"livestream":      MediaSourceContentCategory_LiveStream,
		"meeting":         MediaSourceContentCategory_Meeting,
		"presentation":    MediaSourceContentCategory_Presentation,
		"screenrecording": MediaSourceContentCategory_ScreenRecording,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaSourceContentCategory(input)
	return &out, nil
}
