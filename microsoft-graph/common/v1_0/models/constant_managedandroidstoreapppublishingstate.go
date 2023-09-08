package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidStoreAppPublishingState string

const (
	ManagedAndroidStoreAppPublishingStatenotPublished ManagedAndroidStoreAppPublishingState = "NotPublished"
	ManagedAndroidStoreAppPublishingStateprocessing   ManagedAndroidStoreAppPublishingState = "Processing"
	ManagedAndroidStoreAppPublishingStatepublished    ManagedAndroidStoreAppPublishingState = "Published"
)

func PossibleValuesForManagedAndroidStoreAppPublishingState() []string {
	return []string{
		string(ManagedAndroidStoreAppPublishingStatenotPublished),
		string(ManagedAndroidStoreAppPublishingStateprocessing),
		string(ManagedAndroidStoreAppPublishingStatepublished),
	}
}

func parseManagedAndroidStoreAppPublishingState(input string) (*ManagedAndroidStoreAppPublishingState, error) {
	vals := map[string]ManagedAndroidStoreAppPublishingState{
		"notpublished": ManagedAndroidStoreAppPublishingStatenotPublished,
		"processing":   ManagedAndroidStoreAppPublishingStateprocessing,
		"published":    ManagedAndroidStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidStoreAppPublishingState(input)
	return &out, nil
}
