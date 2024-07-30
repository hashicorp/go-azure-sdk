package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppPublishingState string

const (
	MacOsVppAppPublishingState_NotPublished MacOsVppAppPublishingState = "notPublished"
	MacOsVppAppPublishingState_Processing   MacOsVppAppPublishingState = "processing"
	MacOsVppAppPublishingState_Published    MacOsVppAppPublishingState = "published"
)

func PossibleValuesForMacOsVppAppPublishingState() []string {
	return []string{
		string(MacOsVppAppPublishingState_NotPublished),
		string(MacOsVppAppPublishingState_Processing),
		string(MacOsVppAppPublishingState_Published),
	}
}

func (s *MacOsVppAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOsVppAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOsVppAppPublishingState(input string) (*MacOsVppAppPublishingState, error) {
	vals := map[string]MacOsVppAppPublishingState{
		"notpublished": MacOsVppAppPublishingState_NotPublished,
		"processing":   MacOsVppAppPublishingState_Processing,
		"published":    MacOsVppAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppPublishingState(input)
	return &out, nil
}
