package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobAppPublishingState string

const (
	MacOSLobAppPublishingStatenotPublished MacOSLobAppPublishingState = "NotPublished"
	MacOSLobAppPublishingStateprocessing   MacOSLobAppPublishingState = "Processing"
	MacOSLobAppPublishingStatepublished    MacOSLobAppPublishingState = "Published"
)

func PossibleValuesForMacOSLobAppPublishingState() []string {
	return []string{
		string(MacOSLobAppPublishingStatenotPublished),
		string(MacOSLobAppPublishingStateprocessing),
		string(MacOSLobAppPublishingStatepublished),
	}
}

func parseMacOSLobAppPublishingState(input string) (*MacOSLobAppPublishingState, error) {
	vals := map[string]MacOSLobAppPublishingState{
		"notpublished": MacOSLobAppPublishingStatenotPublished,
		"processing":   MacOSLobAppPublishingStateprocessing,
		"published":    MacOSLobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSLobAppPublishingState(input)
	return &out, nil
}
