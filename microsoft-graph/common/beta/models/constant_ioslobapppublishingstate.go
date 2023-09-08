package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppPublishingState string

const (
	IosLobAppPublishingStatenotPublished IosLobAppPublishingState = "NotPublished"
	IosLobAppPublishingStateprocessing   IosLobAppPublishingState = "Processing"
	IosLobAppPublishingStatepublished    IosLobAppPublishingState = "Published"
)

func PossibleValuesForIosLobAppPublishingState() []string {
	return []string{
		string(IosLobAppPublishingStatenotPublished),
		string(IosLobAppPublishingStateprocessing),
		string(IosLobAppPublishingStatepublished),
	}
}

func parseIosLobAppPublishingState(input string) (*IosLobAppPublishingState, error) {
	vals := map[string]IosLobAppPublishingState{
		"notpublished": IosLobAppPublishingStatenotPublished,
		"processing":   IosLobAppPublishingStateprocessing,
		"published":    IosLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosLobAppPublishingState(input)
	return &out, nil
}
