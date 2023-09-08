package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWebClipPublishingState string

const (
	MacOSWebClipPublishingStatenotPublished MacOSWebClipPublishingState = "NotPublished"
	MacOSWebClipPublishingStateprocessing   MacOSWebClipPublishingState = "Processing"
	MacOSWebClipPublishingStatepublished    MacOSWebClipPublishingState = "Published"
)

func PossibleValuesForMacOSWebClipPublishingState() []string {
	return []string{
		string(MacOSWebClipPublishingStatenotPublished),
		string(MacOSWebClipPublishingStateprocessing),
		string(MacOSWebClipPublishingStatepublished),
	}
}

func parseMacOSWebClipPublishingState(input string) (*MacOSWebClipPublishingState, error) {
	vals := map[string]MacOSWebClipPublishingState{
		"notpublished": MacOSWebClipPublishingStatenotPublished,
		"processing":   MacOSWebClipPublishingStateprocessing,
		"published":    MacOSWebClipPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWebClipPublishingState(input)
	return &out, nil
}
