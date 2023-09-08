package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppPublishingState string

const (
	ManagedAndroidLobAppPublishingStatenotPublished ManagedAndroidLobAppPublishingState = "NotPublished"
	ManagedAndroidLobAppPublishingStateprocessing   ManagedAndroidLobAppPublishingState = "Processing"
	ManagedAndroidLobAppPublishingStatepublished    ManagedAndroidLobAppPublishingState = "Published"
)

func PossibleValuesForManagedAndroidLobAppPublishingState() []string {
	return []string{
		string(ManagedAndroidLobAppPublishingStatenotPublished),
		string(ManagedAndroidLobAppPublishingStateprocessing),
		string(ManagedAndroidLobAppPublishingStatepublished),
	}
}

func parseManagedAndroidLobAppPublishingState(input string) (*ManagedAndroidLobAppPublishingState, error) {
	vals := map[string]ManagedAndroidLobAppPublishingState{
		"notpublished": ManagedAndroidLobAppPublishingStatenotPublished,
		"processing":   ManagedAndroidLobAppPublishingStateprocessing,
		"published":    ManagedAndroidLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppPublishingState(input)
	return &out, nil
}
