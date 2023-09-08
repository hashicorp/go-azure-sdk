package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftDefenderAppPublishingState string

const (
	MacOSMicrosoftDefenderAppPublishingStatenotPublished MacOSMicrosoftDefenderAppPublishingState = "NotPublished"
	MacOSMicrosoftDefenderAppPublishingStateprocessing   MacOSMicrosoftDefenderAppPublishingState = "Processing"
	MacOSMicrosoftDefenderAppPublishingStatepublished    MacOSMicrosoftDefenderAppPublishingState = "Published"
)

func PossibleValuesForMacOSMicrosoftDefenderAppPublishingState() []string {
	return []string{
		string(MacOSMicrosoftDefenderAppPublishingStatenotPublished),
		string(MacOSMicrosoftDefenderAppPublishingStateprocessing),
		string(MacOSMicrosoftDefenderAppPublishingStatepublished),
	}
}

func parseMacOSMicrosoftDefenderAppPublishingState(input string) (*MacOSMicrosoftDefenderAppPublishingState, error) {
	vals := map[string]MacOSMicrosoftDefenderAppPublishingState{
		"notpublished": MacOSMicrosoftDefenderAppPublishingStatenotPublished,
		"processing":   MacOSMicrosoftDefenderAppPublishingStateprocessing,
		"published":    MacOSMicrosoftDefenderAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftDefenderAppPublishingState(input)
	return &out, nil
}
