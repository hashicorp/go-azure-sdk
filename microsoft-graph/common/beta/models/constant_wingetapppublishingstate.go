package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppPublishingState string

const (
	WinGetAppPublishingStatenotPublished WinGetAppPublishingState = "NotPublished"
	WinGetAppPublishingStateprocessing   WinGetAppPublishingState = "Processing"
	WinGetAppPublishingStatepublished    WinGetAppPublishingState = "Published"
)

func PossibleValuesForWinGetAppPublishingState() []string {
	return []string{
		string(WinGetAppPublishingStatenotPublished),
		string(WinGetAppPublishingStateprocessing),
		string(WinGetAppPublishingStatepublished),
	}
}

func parseWinGetAppPublishingState(input string) (*WinGetAppPublishingState, error) {
	vals := map[string]WinGetAppPublishingState{
		"notpublished": WinGetAppPublishingStatenotPublished,
		"processing":   WinGetAppPublishingStateprocessing,
		"published":    WinGetAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WinGetAppPublishingState(input)
	return &out, nil
}
