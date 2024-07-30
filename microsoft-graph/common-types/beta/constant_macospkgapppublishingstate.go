package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkgAppPublishingState string

const (
	MacOSPkgAppPublishingState_NotPublished MacOSPkgAppPublishingState = "notPublished"
	MacOSPkgAppPublishingState_Processing   MacOSPkgAppPublishingState = "processing"
	MacOSPkgAppPublishingState_Published    MacOSPkgAppPublishingState = "published"
)

func PossibleValuesForMacOSPkgAppPublishingState() []string {
	return []string{
		string(MacOSPkgAppPublishingState_NotPublished),
		string(MacOSPkgAppPublishingState_Processing),
		string(MacOSPkgAppPublishingState_Published),
	}
}

func (s *MacOSPkgAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPkgAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPkgAppPublishingState(input string) (*MacOSPkgAppPublishingState, error) {
	vals := map[string]MacOSPkgAppPublishingState{
		"notpublished": MacOSPkgAppPublishingState_NotPublished,
		"processing":   MacOSPkgAppPublishingState_Processing,
		"published":    MacOSPkgAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkgAppPublishingState(input)
	return &out, nil
}
