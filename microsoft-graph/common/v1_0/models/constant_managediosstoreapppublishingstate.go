package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSStoreAppPublishingState string

const (
	ManagedIOSStoreAppPublishingStatenotPublished ManagedIOSStoreAppPublishingState = "NotPublished"
	ManagedIOSStoreAppPublishingStateprocessing   ManagedIOSStoreAppPublishingState = "Processing"
	ManagedIOSStoreAppPublishingStatepublished    ManagedIOSStoreAppPublishingState = "Published"
)

func PossibleValuesForManagedIOSStoreAppPublishingState() []string {
	return []string{
		string(ManagedIOSStoreAppPublishingStatenotPublished),
		string(ManagedIOSStoreAppPublishingStateprocessing),
		string(ManagedIOSStoreAppPublishingStatepublished),
	}
}

func parseManagedIOSStoreAppPublishingState(input string) (*ManagedIOSStoreAppPublishingState, error) {
	vals := map[string]ManagedIOSStoreAppPublishingState{
		"notpublished": ManagedIOSStoreAppPublishingStatenotPublished,
		"processing":   ManagedIOSStoreAppPublishingStateprocessing,
		"published":    ManagedIOSStoreAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSStoreAppPublishingState(input)
	return &out, nil
}
