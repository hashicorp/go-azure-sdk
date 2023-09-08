package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppPublishingState string

const (
	OfficeSuiteAppPublishingStatenotPublished OfficeSuiteAppPublishingState = "NotPublished"
	OfficeSuiteAppPublishingStateprocessing   OfficeSuiteAppPublishingState = "Processing"
	OfficeSuiteAppPublishingStatepublished    OfficeSuiteAppPublishingState = "Published"
)

func PossibleValuesForOfficeSuiteAppPublishingState() []string {
	return []string{
		string(OfficeSuiteAppPublishingStatenotPublished),
		string(OfficeSuiteAppPublishingStateprocessing),
		string(OfficeSuiteAppPublishingStatepublished),
	}
}

func parseOfficeSuiteAppPublishingState(input string) (*OfficeSuiteAppPublishingState, error) {
	vals := map[string]OfficeSuiteAppPublishingState{
		"notpublished": OfficeSuiteAppPublishingStatenotPublished,
		"processing":   OfficeSuiteAppPublishingStateprocessing,
		"published":    OfficeSuiteAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppPublishingState(input)
	return &out, nil
}
