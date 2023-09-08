package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppXPublishingState string

const (
	WindowsAppXPublishingStatenotPublished WindowsAppXPublishingState = "NotPublished"
	WindowsAppXPublishingStateprocessing   WindowsAppXPublishingState = "Processing"
	WindowsAppXPublishingStatepublished    WindowsAppXPublishingState = "Published"
)

func PossibleValuesForWindowsAppXPublishingState() []string {
	return []string{
		string(WindowsAppXPublishingStatenotPublished),
		string(WindowsAppXPublishingStateprocessing),
		string(WindowsAppXPublishingStatepublished),
	}
}

func parseWindowsAppXPublishingState(input string) (*WindowsAppXPublishingState, error) {
	vals := map[string]WindowsAppXPublishingState{
		"notpublished": WindowsAppXPublishingStatenotPublished,
		"processing":   WindowsAppXPublishingStateprocessing,
		"published":    WindowsAppXPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAppXPublishingState(input)
	return &out, nil
}
