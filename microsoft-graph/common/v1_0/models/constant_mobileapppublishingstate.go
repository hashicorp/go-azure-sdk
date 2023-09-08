package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPublishingState string

const (
	MobileAppPublishingStatenotPublished MobileAppPublishingState = "NotPublished"
	MobileAppPublishingStateprocessing   MobileAppPublishingState = "Processing"
	MobileAppPublishingStatepublished    MobileAppPublishingState = "Published"
)

func PossibleValuesForMobileAppPublishingState() []string {
	return []string{
		string(MobileAppPublishingStatenotPublished),
		string(MobileAppPublishingStateprocessing),
		string(MobileAppPublishingStatepublished),
	}
}

func parseMobileAppPublishingState(input string) (*MobileAppPublishingState, error) {
	vals := map[string]MobileAppPublishingState{
		"notpublished": MobileAppPublishingStatenotPublished,
		"processing":   MobileAppPublishingStateprocessing,
		"published":    MobileAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPublishingState(input)
	return &out, nil
}
