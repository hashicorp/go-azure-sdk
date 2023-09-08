package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppPublishingState string

const (
	WebAppPublishingStatenotPublished WebAppPublishingState = "NotPublished"
	WebAppPublishingStateprocessing   WebAppPublishingState = "Processing"
	WebAppPublishingStatepublished    WebAppPublishingState = "Published"
)

func PossibleValuesForWebAppPublishingState() []string {
	return []string{
		string(WebAppPublishingStatenotPublished),
		string(WebAppPublishingStateprocessing),
		string(WebAppPublishingStatepublished),
	}
}

func parseWebAppPublishingState(input string) (*WebAppPublishingState, error) {
	vals := map[string]WebAppPublishingState{
		"notpublished": WebAppPublishingStatenotPublished,
		"processing":   WebAppPublishingStateprocessing,
		"published":    WebAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebAppPublishingState(input)
	return &out, nil
}
