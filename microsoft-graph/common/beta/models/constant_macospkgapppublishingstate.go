package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkgAppPublishingState string

const (
	MacOSPkgAppPublishingStatenotPublished MacOSPkgAppPublishingState = "NotPublished"
	MacOSPkgAppPublishingStateprocessing   MacOSPkgAppPublishingState = "Processing"
	MacOSPkgAppPublishingStatepublished    MacOSPkgAppPublishingState = "Published"
)

func PossibleValuesForMacOSPkgAppPublishingState() []string {
	return []string{
		string(MacOSPkgAppPublishingStatenotPublished),
		string(MacOSPkgAppPublishingStateprocessing),
		string(MacOSPkgAppPublishingStatepublished),
	}
}

func parseMacOSPkgAppPublishingState(input string) (*MacOSPkgAppPublishingState, error) {
	vals := map[string]MacOSPkgAppPublishingState{
		"notpublished": MacOSPkgAppPublishingStatenotPublished,
		"processing":   MacOSPkgAppPublishingStateprocessing,
		"published":    MacOSPkgAppPublishingStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkgAppPublishingState(input)
	return &out, nil
}
