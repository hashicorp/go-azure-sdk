package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDmgAppPublishingState string

const (
	MacOSDmgAppPublishingStatenotPublished MacOSDmgAppPublishingState = "NotPublished"
	MacOSDmgAppPublishingStateprocessing   MacOSDmgAppPublishingState = "Processing"
	MacOSDmgAppPublishingStatepublished    MacOSDmgAppPublishingState = "Published"
)

func PossibleValuesForMacOSDmgAppPublishingState() []string {
	return []string{
		string(MacOSDmgAppPublishingStatenotPublished),
		string(MacOSDmgAppPublishingStateprocessing),
		string(MacOSDmgAppPublishingStatepublished),
	}
}

func parseMacOSDmgAppPublishingState(input string) (*MacOSDmgAppPublishingState, error) {
	vals := map[string]MacOSDmgAppPublishingState{
		"notpublished": MacOSDmgAppPublishingStatenotPublished,
		"processing":   MacOSDmgAppPublishingStateprocessing,
		"published":    MacOSDmgAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDmgAppPublishingState(input)
	return &out, nil
}
