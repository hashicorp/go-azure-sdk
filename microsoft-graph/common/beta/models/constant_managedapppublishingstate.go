package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPublishingState string

const (
	ManagedAppPublishingStatenotPublished ManagedAppPublishingState = "NotPublished"
	ManagedAppPublishingStateprocessing   ManagedAppPublishingState = "Processing"
	ManagedAppPublishingStatepublished    ManagedAppPublishingState = "Published"
)

func PossibleValuesForManagedAppPublishingState() []string {
	return []string{
		string(ManagedAppPublishingStatenotPublished),
		string(ManagedAppPublishingStateprocessing),
		string(ManagedAppPublishingStatepublished),
	}
}

func parseManagedAppPublishingState(input string) (*ManagedAppPublishingState, error) {
	vals := map[string]ManagedAppPublishingState{
		"notpublished": ManagedAppPublishingStatenotPublished,
		"processing":   ManagedAppPublishingStateprocessing,
		"published":    ManagedAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppPublishingState(input)
	return &out, nil
}
