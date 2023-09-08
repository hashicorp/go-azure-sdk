package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81StoreAppPublishingState string

const (
	WindowsPhone81StoreAppPublishingStatenotPublished WindowsPhone81StoreAppPublishingState = "NotPublished"
	WindowsPhone81StoreAppPublishingStateprocessing   WindowsPhone81StoreAppPublishingState = "Processing"
	WindowsPhone81StoreAppPublishingStatepublished    WindowsPhone81StoreAppPublishingState = "Published"
)

func PossibleValuesForWindowsPhone81StoreAppPublishingState() []string {
	return []string{
		string(WindowsPhone81StoreAppPublishingStatenotPublished),
		string(WindowsPhone81StoreAppPublishingStateprocessing),
		string(WindowsPhone81StoreAppPublishingStatepublished),
	}
}

func parseWindowsPhone81StoreAppPublishingState(input string) (*WindowsPhone81StoreAppPublishingState, error) {
	vals := map[string]WindowsPhone81StoreAppPublishingState{
		"notpublished": WindowsPhone81StoreAppPublishingStatenotPublished,
		"processing":   WindowsPhone81StoreAppPublishingStateprocessing,
		"published":    WindowsPhone81StoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81StoreAppPublishingState(input)
	return &out, nil
}
