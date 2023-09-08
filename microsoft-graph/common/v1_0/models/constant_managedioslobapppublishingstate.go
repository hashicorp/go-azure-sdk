package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSLobAppPublishingState string

const (
	ManagedIOSLobAppPublishingStatenotPublished ManagedIOSLobAppPublishingState = "NotPublished"
	ManagedIOSLobAppPublishingStateprocessing   ManagedIOSLobAppPublishingState = "Processing"
	ManagedIOSLobAppPublishingStatepublished    ManagedIOSLobAppPublishingState = "Published"
)

func PossibleValuesForManagedIOSLobAppPublishingState() []string {
	return []string{
		string(ManagedIOSLobAppPublishingStatenotPublished),
		string(ManagedIOSLobAppPublishingStateprocessing),
		string(ManagedIOSLobAppPublishingStatepublished),
	}
}

func parseManagedIOSLobAppPublishingState(input string) (*ManagedIOSLobAppPublishingState, error) {
	vals := map[string]ManagedIOSLobAppPublishingState{
		"notpublished": ManagedIOSLobAppPublishingStatenotPublished,
		"processing":   ManagedIOSLobAppPublishingStateprocessing,
		"published":    ManagedIOSLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSLobAppPublishingState(input)
	return &out, nil
}
