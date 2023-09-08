package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidLobAppPublishingState string

const (
	AndroidLobAppPublishingStatenotPublished AndroidLobAppPublishingState = "NotPublished"
	AndroidLobAppPublishingStateprocessing   AndroidLobAppPublishingState = "Processing"
	AndroidLobAppPublishingStatepublished    AndroidLobAppPublishingState = "Published"
)

func PossibleValuesForAndroidLobAppPublishingState() []string {
	return []string{
		string(AndroidLobAppPublishingStatenotPublished),
		string(AndroidLobAppPublishingStateprocessing),
		string(AndroidLobAppPublishingStatepublished),
	}
}

func parseAndroidLobAppPublishingState(input string) (*AndroidLobAppPublishingState, error) {
	vals := map[string]AndroidLobAppPublishingState{
		"notpublished": AndroidLobAppPublishingStatenotPublished,
		"processing":   AndroidLobAppPublishingStateprocessing,
		"published":    AndroidLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidLobAppPublishingState(input)
	return &out, nil
}
