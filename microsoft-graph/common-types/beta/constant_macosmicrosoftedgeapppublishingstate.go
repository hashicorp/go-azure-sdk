package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftEdgeAppPublishingState string

const (
	MacOSMicrosoftEdgeAppPublishingState_NotPublished MacOSMicrosoftEdgeAppPublishingState = "notPublished"
	MacOSMicrosoftEdgeAppPublishingState_Processing   MacOSMicrosoftEdgeAppPublishingState = "processing"
	MacOSMicrosoftEdgeAppPublishingState_Published    MacOSMicrosoftEdgeAppPublishingState = "published"
)

func PossibleValuesForMacOSMicrosoftEdgeAppPublishingState() []string {
	return []string{
		string(MacOSMicrosoftEdgeAppPublishingState_NotPublished),
		string(MacOSMicrosoftEdgeAppPublishingState_Processing),
		string(MacOSMicrosoftEdgeAppPublishingState_Published),
	}
}

func (s *MacOSMicrosoftEdgeAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSMicrosoftEdgeAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSMicrosoftEdgeAppPublishingState(input string) (*MacOSMicrosoftEdgeAppPublishingState, error) {
	vals := map[string]MacOSMicrosoftEdgeAppPublishingState{
		"notpublished": MacOSMicrosoftEdgeAppPublishingState_NotPublished,
		"processing":   MacOSMicrosoftEdgeAppPublishingState_Processing,
		"published":    MacOSMicrosoftEdgeAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftEdgeAppPublishingState(input)
	return &out, nil
}
