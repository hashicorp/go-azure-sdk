package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreWebAppPublishingState string

const (
	AndroidManagedStoreWebAppPublishingStatenotPublished AndroidManagedStoreWebAppPublishingState = "NotPublished"
	AndroidManagedStoreWebAppPublishingStateprocessing   AndroidManagedStoreWebAppPublishingState = "Processing"
	AndroidManagedStoreWebAppPublishingStatepublished    AndroidManagedStoreWebAppPublishingState = "Published"
)

func PossibleValuesForAndroidManagedStoreWebAppPublishingState() []string {
	return []string{
		string(AndroidManagedStoreWebAppPublishingStatenotPublished),
		string(AndroidManagedStoreWebAppPublishingStateprocessing),
		string(AndroidManagedStoreWebAppPublishingStatepublished),
	}
}

func parseAndroidManagedStoreWebAppPublishingState(input string) (*AndroidManagedStoreWebAppPublishingState, error) {
	vals := map[string]AndroidManagedStoreWebAppPublishingState{
		"notpublished": AndroidManagedStoreWebAppPublishingStatenotPublished,
		"processing":   AndroidManagedStoreWebAppPublishingStateprocessing,
		"published":    AndroidManagedStoreWebAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreWebAppPublishingState(input)
	return &out, nil
}
