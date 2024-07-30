package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftDefenderAppPublishingState string

const (
	MacOSMicrosoftDefenderAppPublishingState_NotPublished MacOSMicrosoftDefenderAppPublishingState = "notPublished"
	MacOSMicrosoftDefenderAppPublishingState_Processing   MacOSMicrosoftDefenderAppPublishingState = "processing"
	MacOSMicrosoftDefenderAppPublishingState_Published    MacOSMicrosoftDefenderAppPublishingState = "published"
)

func PossibleValuesForMacOSMicrosoftDefenderAppPublishingState() []string {
	return []string{
		string(MacOSMicrosoftDefenderAppPublishingState_NotPublished),
		string(MacOSMicrosoftDefenderAppPublishingState_Processing),
		string(MacOSMicrosoftDefenderAppPublishingState_Published),
	}
}

func (s *MacOSMicrosoftDefenderAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSMicrosoftDefenderAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSMicrosoftDefenderAppPublishingState(input string) (*MacOSMicrosoftDefenderAppPublishingState, error) {
	vals := map[string]MacOSMicrosoftDefenderAppPublishingState{
		"notpublished": MacOSMicrosoftDefenderAppPublishingState_NotPublished,
		"processing":   MacOSMicrosoftDefenderAppPublishingState_Processing,
		"published":    MacOSMicrosoftDefenderAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftDefenderAppPublishingState(input)
	return &out, nil
}
