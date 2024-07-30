package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppPublishingState string

const (
	OfficeSuiteAppPublishingState_NotPublished OfficeSuiteAppPublishingState = "notPublished"
	OfficeSuiteAppPublishingState_Processing   OfficeSuiteAppPublishingState = "processing"
	OfficeSuiteAppPublishingState_Published    OfficeSuiteAppPublishingState = "published"
)

func PossibleValuesForOfficeSuiteAppPublishingState() []string {
	return []string{
		string(OfficeSuiteAppPublishingState_NotPublished),
		string(OfficeSuiteAppPublishingState_Processing),
		string(OfficeSuiteAppPublishingState_Published),
	}
}

func (s *OfficeSuiteAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppPublishingState(input string) (*OfficeSuiteAppPublishingState, error) {
	vals := map[string]OfficeSuiteAppPublishingState{
		"notpublished": OfficeSuiteAppPublishingState_NotPublished,
		"processing":   OfficeSuiteAppPublishingState_Processing,
		"published":    OfficeSuiteAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppPublishingState(input)
	return &out, nil
}
