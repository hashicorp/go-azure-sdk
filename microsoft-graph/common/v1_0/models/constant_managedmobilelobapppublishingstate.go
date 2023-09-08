package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileLobAppPublishingState string

const (
	ManagedMobileLobAppPublishingStatenotPublished ManagedMobileLobAppPublishingState = "NotPublished"
	ManagedMobileLobAppPublishingStateprocessing   ManagedMobileLobAppPublishingState = "Processing"
	ManagedMobileLobAppPublishingStatepublished    ManagedMobileLobAppPublishingState = "Published"
)

func PossibleValuesForManagedMobileLobAppPublishingState() []string {
	return []string{
		string(ManagedMobileLobAppPublishingStatenotPublished),
		string(ManagedMobileLobAppPublishingStateprocessing),
		string(ManagedMobileLobAppPublishingStatepublished),
	}
}

func parseManagedMobileLobAppPublishingState(input string) (*ManagedMobileLobAppPublishingState, error) {
	vals := map[string]ManagedMobileLobAppPublishingState{
		"notpublished": ManagedMobileLobAppPublishingStatenotPublished,
		"processing":   ManagedMobileLobAppPublishingStateprocessing,
		"published":    ManagedMobileLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedMobileLobAppPublishingState(input)
	return &out, nil
}
