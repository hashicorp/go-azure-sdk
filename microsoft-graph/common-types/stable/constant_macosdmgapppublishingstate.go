package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDmgAppPublishingState string

const (
	MacOSDmgAppPublishingState_NotPublished MacOSDmgAppPublishingState = "notPublished"
	MacOSDmgAppPublishingState_Processing   MacOSDmgAppPublishingState = "processing"
	MacOSDmgAppPublishingState_Published    MacOSDmgAppPublishingState = "published"
)

func PossibleValuesForMacOSDmgAppPublishingState() []string {
	return []string{
		string(MacOSDmgAppPublishingState_NotPublished),
		string(MacOSDmgAppPublishingState_Processing),
		string(MacOSDmgAppPublishingState_Published),
	}
}

func (s *MacOSDmgAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSDmgAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSDmgAppPublishingState(input string) (*MacOSDmgAppPublishingState, error) {
	vals := map[string]MacOSDmgAppPublishingState{
		"notpublished": MacOSDmgAppPublishingState_NotPublished,
		"processing":   MacOSDmgAppPublishingState_Processing,
		"published":    MacOSDmgAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDmgAppPublishingState(input)
	return &out, nil
}
