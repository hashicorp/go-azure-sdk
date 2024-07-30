package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobAppPublishingState string

const (
	MacOSLobAppPublishingState_NotPublished MacOSLobAppPublishingState = "notPublished"
	MacOSLobAppPublishingState_Processing   MacOSLobAppPublishingState = "processing"
	MacOSLobAppPublishingState_Published    MacOSLobAppPublishingState = "published"
)

func PossibleValuesForMacOSLobAppPublishingState() []string {
	return []string{
		string(MacOSLobAppPublishingState_NotPublished),
		string(MacOSLobAppPublishingState_Processing),
		string(MacOSLobAppPublishingState_Published),
	}
}

func (s *MacOSLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSLobAppPublishingState(input string) (*MacOSLobAppPublishingState, error) {
	vals := map[string]MacOSLobAppPublishingState{
		"notpublished": MacOSLobAppPublishingState_NotPublished,
		"processing":   MacOSLobAppPublishingState_Processing,
		"published":    MacOSLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSLobAppPublishingState(input)
	return &out, nil
}
