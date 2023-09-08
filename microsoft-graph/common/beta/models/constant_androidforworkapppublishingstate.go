package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppPublishingState string

const (
	AndroidForWorkAppPublishingStatenotPublished AndroidForWorkAppPublishingState = "NotPublished"
	AndroidForWorkAppPublishingStateprocessing   AndroidForWorkAppPublishingState = "Processing"
	AndroidForWorkAppPublishingStatepublished    AndroidForWorkAppPublishingState = "Published"
)

func PossibleValuesForAndroidForWorkAppPublishingState() []string {
	return []string{
		string(AndroidForWorkAppPublishingStatenotPublished),
		string(AndroidForWorkAppPublishingStateprocessing),
		string(AndroidForWorkAppPublishingStatepublished),
	}
}

func parseAndroidForWorkAppPublishingState(input string) (*AndroidForWorkAppPublishingState, error) {
	vals := map[string]AndroidForWorkAppPublishingState{
		"notpublished": AndroidForWorkAppPublishingStatenotPublished,
		"processing":   AndroidForWorkAppPublishingStateprocessing,
		"published":    AndroidForWorkAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkAppPublishingState(input)
	return &out, nil
}
