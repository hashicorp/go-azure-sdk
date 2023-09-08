package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosiPadOSWebClipPublishingState string

const (
	IosiPadOSWebClipPublishingStatenotPublished IosiPadOSWebClipPublishingState = "NotPublished"
	IosiPadOSWebClipPublishingStateprocessing   IosiPadOSWebClipPublishingState = "Processing"
	IosiPadOSWebClipPublishingStatepublished    IosiPadOSWebClipPublishingState = "Published"
)

func PossibleValuesForIosiPadOSWebClipPublishingState() []string {
	return []string{
		string(IosiPadOSWebClipPublishingStatenotPublished),
		string(IosiPadOSWebClipPublishingStateprocessing),
		string(IosiPadOSWebClipPublishingStatepublished),
	}
}

func parseIosiPadOSWebClipPublishingState(input string) (*IosiPadOSWebClipPublishingState, error) {
	vals := map[string]IosiPadOSWebClipPublishingState{
		"notpublished": IosiPadOSWebClipPublishingStatenotPublished,
		"processing":   IosiPadOSWebClipPublishingStateprocessing,
		"published":    IosiPadOSWebClipPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosiPadOSWebClipPublishingState(input)
	return &out, nil
}
