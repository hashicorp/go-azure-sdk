package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSLobAppPublishingState string

const (
	ManagedIOSLobAppPublishingState_NotPublished ManagedIOSLobAppPublishingState = "notPublished"
	ManagedIOSLobAppPublishingState_Processing   ManagedIOSLobAppPublishingState = "processing"
	ManagedIOSLobAppPublishingState_Published    ManagedIOSLobAppPublishingState = "published"
)

func PossibleValuesForManagedIOSLobAppPublishingState() []string {
	return []string{
		string(ManagedIOSLobAppPublishingState_NotPublished),
		string(ManagedIOSLobAppPublishingState_Processing),
		string(ManagedIOSLobAppPublishingState_Published),
	}
}

func (s *ManagedIOSLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedIOSLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedIOSLobAppPublishingState(input string) (*ManagedIOSLobAppPublishingState, error) {
	vals := map[string]ManagedIOSLobAppPublishingState{
		"notpublished": ManagedIOSLobAppPublishingState_NotPublished,
		"processing":   ManagedIOSLobAppPublishingState_Processing,
		"published":    ManagedIOSLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSLobAppPublishingState(input)
	return &out, nil
}
