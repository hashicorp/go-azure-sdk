package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXPublishingState string

const (
	WindowsUniversalAppXPublishingStatenotPublished WindowsUniversalAppXPublishingState = "NotPublished"
	WindowsUniversalAppXPublishingStateprocessing   WindowsUniversalAppXPublishingState = "Processing"
	WindowsUniversalAppXPublishingStatepublished    WindowsUniversalAppXPublishingState = "Published"
)

func PossibleValuesForWindowsUniversalAppXPublishingState() []string {
	return []string{
		string(WindowsUniversalAppXPublishingStatenotPublished),
		string(WindowsUniversalAppXPublishingStateprocessing),
		string(WindowsUniversalAppXPublishingStatepublished),
	}
}

func parseWindowsUniversalAppXPublishingState(input string) (*WindowsUniversalAppXPublishingState, error) {
	vals := map[string]WindowsUniversalAppXPublishingState{
		"notpublished": WindowsUniversalAppXPublishingStatenotPublished,
		"processing":   WindowsUniversalAppXPublishingStateprocessing,
		"published":    WindowsUniversalAppXPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXPublishingState(input)
	return &out, nil
}
