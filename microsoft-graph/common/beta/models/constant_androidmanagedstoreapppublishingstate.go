package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppPublishingState string

const (
	AndroidManagedStoreAppPublishingStatenotPublished AndroidManagedStoreAppPublishingState = "NotPublished"
	AndroidManagedStoreAppPublishingStateprocessing   AndroidManagedStoreAppPublishingState = "Processing"
	AndroidManagedStoreAppPublishingStatepublished    AndroidManagedStoreAppPublishingState = "Published"
)

func PossibleValuesForAndroidManagedStoreAppPublishingState() []string {
	return []string{
		string(AndroidManagedStoreAppPublishingStatenotPublished),
		string(AndroidManagedStoreAppPublishingStateprocessing),
		string(AndroidManagedStoreAppPublishingStatepublished),
	}
}

func parseAndroidManagedStoreAppPublishingState(input string) (*AndroidManagedStoreAppPublishingState, error) {
	vals := map[string]AndroidManagedStoreAppPublishingState{
		"notpublished": AndroidManagedStoreAppPublishingStatenotPublished,
		"processing":   AndroidManagedStoreAppPublishingStateprocessing,
		"published":    AndroidManagedStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppPublishingState(input)
	return &out, nil
}
