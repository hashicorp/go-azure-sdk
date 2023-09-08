package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXBundlePublishingState string

const (
	WindowsPhone81AppXBundlePublishingStatenotPublished WindowsPhone81AppXBundlePublishingState = "NotPublished"
	WindowsPhone81AppXBundlePublishingStateprocessing   WindowsPhone81AppXBundlePublishingState = "Processing"
	WindowsPhone81AppXBundlePublishingStatepublished    WindowsPhone81AppXBundlePublishingState = "Published"
)

func PossibleValuesForWindowsPhone81AppXBundlePublishingState() []string {
	return []string{
		string(WindowsPhone81AppXBundlePublishingStatenotPublished),
		string(WindowsPhone81AppXBundlePublishingStateprocessing),
		string(WindowsPhone81AppXBundlePublishingStatepublished),
	}
}

func parseWindowsPhone81AppXBundlePublishingState(input string) (*WindowsPhone81AppXBundlePublishingState, error) {
	vals := map[string]WindowsPhone81AppXBundlePublishingState{
		"notpublished": WindowsPhone81AppXBundlePublishingStatenotPublished,
		"processing":   WindowsPhone81AppXBundlePublishingStateprocessing,
		"published":    WindowsPhone81AppXBundlePublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXBundlePublishingState(input)
	return &out, nil
}
