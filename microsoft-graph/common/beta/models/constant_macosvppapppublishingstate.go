package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppPublishingState string

const (
	MacOsVppAppPublishingStatenotPublished MacOsVppAppPublishingState = "NotPublished"
	MacOsVppAppPublishingStateprocessing   MacOsVppAppPublishingState = "Processing"
	MacOsVppAppPublishingStatepublished    MacOsVppAppPublishingState = "Published"
)

func PossibleValuesForMacOsVppAppPublishingState() []string {
	return []string{
		string(MacOsVppAppPublishingStatenotPublished),
		string(MacOsVppAppPublishingStateprocessing),
		string(MacOsVppAppPublishingStatepublished),
	}
}

func parseMacOsVppAppPublishingState(input string) (*MacOsVppAppPublishingState, error) {
	vals := map[string]MacOsVppAppPublishingState{
		"notpublished": MacOsVppAppPublishingStatenotPublished,
		"processing":   MacOsVppAppPublishingStateprocessing,
		"published":    MacOsVppAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppPublishingState(input)
	return &out, nil
}
