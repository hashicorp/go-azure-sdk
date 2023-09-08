package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaSourceContentCategory string

const (
	MediaSourceContentCategoryliveStream      MediaSourceContentCategory = "LiveStream"
	MediaSourceContentCategorymeeting         MediaSourceContentCategory = "Meeting"
	MediaSourceContentCategorypresentation    MediaSourceContentCategory = "Presentation"
	MediaSourceContentCategoryscreenRecording MediaSourceContentCategory = "ScreenRecording"
)

func PossibleValuesForMediaSourceContentCategory() []string {
	return []string{
		string(MediaSourceContentCategoryliveStream),
		string(MediaSourceContentCategorymeeting),
		string(MediaSourceContentCategorypresentation),
		string(MediaSourceContentCategoryscreenRecording),
	}
}

func parseMediaSourceContentCategory(input string) (*MediaSourceContentCategory, error) {
	vals := map[string]MediaSourceContentCategory{
		"livestream":      MediaSourceContentCategoryliveStream,
		"meeting":         MediaSourceContentCategorymeeting,
		"presentation":    MediaSourceContentCategorypresentation,
		"screenrecording": MediaSourceContentCategoryscreenRecording,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaSourceContentCategory(input)
	return &out, nil
}
