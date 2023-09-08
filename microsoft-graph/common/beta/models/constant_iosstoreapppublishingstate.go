package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosStoreAppPublishingState string

const (
	IosStoreAppPublishingStatenotPublished IosStoreAppPublishingState = "NotPublished"
	IosStoreAppPublishingStateprocessing   IosStoreAppPublishingState = "Processing"
	IosStoreAppPublishingStatepublished    IosStoreAppPublishingState = "Published"
)

func PossibleValuesForIosStoreAppPublishingState() []string {
	return []string{
		string(IosStoreAppPublishingStatenotPublished),
		string(IosStoreAppPublishingStateprocessing),
		string(IosStoreAppPublishingStatepublished),
	}
}

func parseIosStoreAppPublishingState(input string) (*IosStoreAppPublishingState, error) {
	vals := map[string]IosStoreAppPublishingState{
		"notpublished": IosStoreAppPublishingStatenotPublished,
		"processing":   IosStoreAppPublishingStateprocessing,
		"published":    IosStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosStoreAppPublishingState(input)
	return &out, nil
}
