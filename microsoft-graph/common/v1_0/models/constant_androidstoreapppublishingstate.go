package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidStoreAppPublishingState string

const (
	AndroidStoreAppPublishingStatenotPublished AndroidStoreAppPublishingState = "NotPublished"
	AndroidStoreAppPublishingStateprocessing   AndroidStoreAppPublishingState = "Processing"
	AndroidStoreAppPublishingStatepublished    AndroidStoreAppPublishingState = "Published"
)

func PossibleValuesForAndroidStoreAppPublishingState() []string {
	return []string{
		string(AndroidStoreAppPublishingStatenotPublished),
		string(AndroidStoreAppPublishingStateprocessing),
		string(AndroidStoreAppPublishingStatepublished),
	}
}

func parseAndroidStoreAppPublishingState(input string) (*AndroidStoreAppPublishingState, error) {
	vals := map[string]AndroidStoreAppPublishingState{
		"notpublished": AndroidStoreAppPublishingStatenotPublished,
		"processing":   AndroidStoreAppPublishingStateprocessing,
		"published":    AndroidStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidStoreAppPublishingState(input)
	return &out, nil
}
