package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPublishingState string

const (
	Win32LobAppPublishingStatenotPublished Win32LobAppPublishingState = "NotPublished"
	Win32LobAppPublishingStateprocessing   Win32LobAppPublishingState = "Processing"
	Win32LobAppPublishingStatepublished    Win32LobAppPublishingState = "Published"
)

func PossibleValuesForWin32LobAppPublishingState() []string {
	return []string{
		string(Win32LobAppPublishingStatenotPublished),
		string(Win32LobAppPublishingStateprocessing),
		string(Win32LobAppPublishingStatepublished),
	}
}

func parseWin32LobAppPublishingState(input string) (*Win32LobAppPublishingState, error) {
	vals := map[string]Win32LobAppPublishingState{
		"notpublished": Win32LobAppPublishingStatenotPublished,
		"processing":   Win32LobAppPublishingStateprocessing,
		"published":    Win32LobAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPublishingState(input)
	return &out, nil
}
