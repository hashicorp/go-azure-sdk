package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMicrosoftEdgeAppPublishingState string

const (
	WindowsMicrosoftEdgeAppPublishingStatenotPublished WindowsMicrosoftEdgeAppPublishingState = "NotPublished"
	WindowsMicrosoftEdgeAppPublishingStateprocessing   WindowsMicrosoftEdgeAppPublishingState = "Processing"
	WindowsMicrosoftEdgeAppPublishingStatepublished    WindowsMicrosoftEdgeAppPublishingState = "Published"
)

func PossibleValuesForWindowsMicrosoftEdgeAppPublishingState() []string {
	return []string{
		string(WindowsMicrosoftEdgeAppPublishingStatenotPublished),
		string(WindowsMicrosoftEdgeAppPublishingStateprocessing),
		string(WindowsMicrosoftEdgeAppPublishingStatepublished),
	}
}

func parseWindowsMicrosoftEdgeAppPublishingState(input string) (*WindowsMicrosoftEdgeAppPublishingState, error) {
	vals := map[string]WindowsMicrosoftEdgeAppPublishingState{
		"notpublished": WindowsMicrosoftEdgeAppPublishingStatenotPublished,
		"processing":   WindowsMicrosoftEdgeAppPublishingStateprocessing,
		"published":    WindowsMicrosoftEdgeAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMicrosoftEdgeAppPublishingState(input)
	return &out, nil
}
