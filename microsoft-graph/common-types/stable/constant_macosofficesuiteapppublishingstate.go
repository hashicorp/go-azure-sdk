package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSOfficeSuiteAppPublishingState string

const (
	MacOSOfficeSuiteAppPublishingState_NotPublished MacOSOfficeSuiteAppPublishingState = "notPublished"
	MacOSOfficeSuiteAppPublishingState_Processing   MacOSOfficeSuiteAppPublishingState = "processing"
	MacOSOfficeSuiteAppPublishingState_Published    MacOSOfficeSuiteAppPublishingState = "published"
)

func PossibleValuesForMacOSOfficeSuiteAppPublishingState() []string {
	return []string{
		string(MacOSOfficeSuiteAppPublishingState_NotPublished),
		string(MacOSOfficeSuiteAppPublishingState_Processing),
		string(MacOSOfficeSuiteAppPublishingState_Published),
	}
}

func (s *MacOSOfficeSuiteAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSOfficeSuiteAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSOfficeSuiteAppPublishingState(input string) (*MacOSOfficeSuiteAppPublishingState, error) {
	vals := map[string]MacOSOfficeSuiteAppPublishingState{
		"notpublished": MacOSOfficeSuiteAppPublishingState_NotPublished,
		"processing":   MacOSOfficeSuiteAppPublishingState_Processing,
		"published":    MacOSOfficeSuiteAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSOfficeSuiteAppPublishingState(input)
	return &out, nil
}
