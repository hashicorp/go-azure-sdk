package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMobileMSIPublishingState string

const (
	WindowsMobileMSIPublishingStatenotPublished WindowsMobileMSIPublishingState = "NotPublished"
	WindowsMobileMSIPublishingStateprocessing   WindowsMobileMSIPublishingState = "Processing"
	WindowsMobileMSIPublishingStatepublished    WindowsMobileMSIPublishingState = "Published"
)

func PossibleValuesForWindowsMobileMSIPublishingState() []string {
	return []string{
		string(WindowsMobileMSIPublishingStatenotPublished),
		string(WindowsMobileMSIPublishingStateprocessing),
		string(WindowsMobileMSIPublishingStatepublished),
	}
}

func parseWindowsMobileMSIPublishingState(input string) (*WindowsMobileMSIPublishingState, error) {
	vals := map[string]WindowsMobileMSIPublishingState{
		"notpublished": WindowsMobileMSIPublishingStatenotPublished,
		"processing":   WindowsMobileMSIPublishingStateprocessing,
		"published":    WindowsMobileMSIPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMobileMSIPublishingState(input)
	return &out, nil
}
