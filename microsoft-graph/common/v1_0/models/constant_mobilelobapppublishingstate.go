package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileLobAppPublishingState string

const (
	MobileLobAppPublishingStatenotPublished MobileLobAppPublishingState = "NotPublished"
	MobileLobAppPublishingStateprocessing   MobileLobAppPublishingState = "Processing"
	MobileLobAppPublishingStatepublished    MobileLobAppPublishingState = "Published"
)

func PossibleValuesForMobileLobAppPublishingState() []string {
	return []string{
		string(MobileLobAppPublishingStatenotPublished),
		string(MobileLobAppPublishingStateprocessing),
		string(MobileLobAppPublishingStatepublished),
	}
}

func parseMobileLobAppPublishingState(input string) (*MobileLobAppPublishingState, error) {
	vals := map[string]MobileLobAppPublishingState{
		"notpublished": MobileLobAppPublishingStatenotPublished,
		"processing":   MobileLobAppPublishingStateprocessing,
		"published":    MobileLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileLobAppPublishingState(input)
	return &out, nil
}
