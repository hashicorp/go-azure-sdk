package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSOfficeSuiteAppPublishingState string

const (
	MacOSOfficeSuiteAppPublishingStatenotPublished MacOSOfficeSuiteAppPublishingState = "NotPublished"
	MacOSOfficeSuiteAppPublishingStateprocessing   MacOSOfficeSuiteAppPublishingState = "Processing"
	MacOSOfficeSuiteAppPublishingStatepublished    MacOSOfficeSuiteAppPublishingState = "Published"
)

func PossibleValuesForMacOSOfficeSuiteAppPublishingState() []string {
	return []string{
		string(MacOSOfficeSuiteAppPublishingStatenotPublished),
		string(MacOSOfficeSuiteAppPublishingStateprocessing),
		string(MacOSOfficeSuiteAppPublishingStatepublished),
	}
}

func parseMacOSOfficeSuiteAppPublishingState(input string) (*MacOSOfficeSuiteAppPublishingState, error) {
	vals := map[string]MacOSOfficeSuiteAppPublishingState{
		"notpublished": MacOSOfficeSuiteAppPublishingStatenotPublished,
		"processing":   MacOSOfficeSuiteAppPublishingStateprocessing,
		"published":    MacOSOfficeSuiteAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSOfficeSuiteAppPublishingState(input)
	return &out, nil
}
