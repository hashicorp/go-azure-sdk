package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppPublishingState string

const (
	IosVppAppPublishingStatenotPublished IosVppAppPublishingState = "NotPublished"
	IosVppAppPublishingStateprocessing   IosVppAppPublishingState = "Processing"
	IosVppAppPublishingStatepublished    IosVppAppPublishingState = "Published"
)

func PossibleValuesForIosVppAppPublishingState() []string {
	return []string{
		string(IosVppAppPublishingStatenotPublished),
		string(IosVppAppPublishingStateprocessing),
		string(IosVppAppPublishingStatepublished),
	}
}

func parseIosVppAppPublishingState(input string) (*IosVppAppPublishingState, error) {
	vals := map[string]IosVppAppPublishingState{
		"notpublished": IosVppAppPublishingStatenotPublished,
		"processing":   IosVppAppPublishingStateprocessing,
		"published":    IosVppAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppPublishingState(input)
	return &out, nil
}
