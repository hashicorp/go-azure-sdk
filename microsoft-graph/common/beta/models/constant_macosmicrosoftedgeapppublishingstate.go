package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftEdgeAppPublishingState string

const (
	MacOSMicrosoftEdgeAppPublishingStatenotPublished MacOSMicrosoftEdgeAppPublishingState = "NotPublished"
	MacOSMicrosoftEdgeAppPublishingStateprocessing   MacOSMicrosoftEdgeAppPublishingState = "Processing"
	MacOSMicrosoftEdgeAppPublishingStatepublished    MacOSMicrosoftEdgeAppPublishingState = "Published"
)

func PossibleValuesForMacOSMicrosoftEdgeAppPublishingState() []string {
	return []string{
		string(MacOSMicrosoftEdgeAppPublishingStatenotPublished),
		string(MacOSMicrosoftEdgeAppPublishingStateprocessing),
		string(MacOSMicrosoftEdgeAppPublishingStatepublished),
	}
}

func parseMacOSMicrosoftEdgeAppPublishingState(input string) (*MacOSMicrosoftEdgeAppPublishingState, error) {
	vals := map[string]MacOSMicrosoftEdgeAppPublishingState{
		"notpublished": MacOSMicrosoftEdgeAppPublishingStatenotPublished,
		"processing":   MacOSMicrosoftEdgeAppPublishingStateprocessing,
		"published":    MacOSMicrosoftEdgeAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftEdgeAppPublishingState(input)
	return &out, nil
}
