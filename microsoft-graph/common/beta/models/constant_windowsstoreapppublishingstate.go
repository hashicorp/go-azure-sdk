package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsStoreAppPublishingState string

const (
	WindowsStoreAppPublishingStatenotPublished WindowsStoreAppPublishingState = "NotPublished"
	WindowsStoreAppPublishingStateprocessing   WindowsStoreAppPublishingState = "Processing"
	WindowsStoreAppPublishingStatepublished    WindowsStoreAppPublishingState = "Published"
)

func PossibleValuesForWindowsStoreAppPublishingState() []string {
	return []string{
		string(WindowsStoreAppPublishingStatenotPublished),
		string(WindowsStoreAppPublishingStateprocessing),
		string(WindowsStoreAppPublishingStatepublished),
	}
}

func parseWindowsStoreAppPublishingState(input string) (*WindowsStoreAppPublishingState, error) {
	vals := map[string]WindowsStoreAppPublishingState{
		"notpublished": WindowsStoreAppPublishingStatenotPublished,
		"processing":   WindowsStoreAppPublishingStateprocessing,
		"published":    WindowsStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsStoreAppPublishingState(input)
	return &out, nil
}
