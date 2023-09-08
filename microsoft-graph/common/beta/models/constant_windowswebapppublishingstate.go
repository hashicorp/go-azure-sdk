package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWebAppPublishingState string

const (
	WindowsWebAppPublishingStatenotPublished WindowsWebAppPublishingState = "NotPublished"
	WindowsWebAppPublishingStateprocessing   WindowsWebAppPublishingState = "Processing"
	WindowsWebAppPublishingStatepublished    WindowsWebAppPublishingState = "Published"
)

func PossibleValuesForWindowsWebAppPublishingState() []string {
	return []string{
		string(WindowsWebAppPublishingStatenotPublished),
		string(WindowsWebAppPublishingStateprocessing),
		string(WindowsWebAppPublishingStatepublished),
	}
}

func parseWindowsWebAppPublishingState(input string) (*WindowsWebAppPublishingState, error) {
	vals := map[string]WindowsWebAppPublishingState{
		"notpublished": WindowsWebAppPublishingStatenotPublished,
		"processing":   WindowsWebAppPublishingStateprocessing,
		"published":    WindowsWebAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWebAppPublishingState(input)
	return &out, nil
}
