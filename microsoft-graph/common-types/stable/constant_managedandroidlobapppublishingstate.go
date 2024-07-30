package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppPublishingState string

const (
	ManagedAndroidLobAppPublishingState_NotPublished ManagedAndroidLobAppPublishingState = "notPublished"
	ManagedAndroidLobAppPublishingState_Processing   ManagedAndroidLobAppPublishingState = "processing"
	ManagedAndroidLobAppPublishingState_Published    ManagedAndroidLobAppPublishingState = "published"
)

func PossibleValuesForManagedAndroidLobAppPublishingState() []string {
	return []string{
		string(ManagedAndroidLobAppPublishingState_NotPublished),
		string(ManagedAndroidLobAppPublishingState_Processing),
		string(ManagedAndroidLobAppPublishingState_Published),
	}
}

func (s *ManagedAndroidLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAndroidLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAndroidLobAppPublishingState(input string) (*ManagedAndroidLobAppPublishingState, error) {
	vals := map[string]ManagedAndroidLobAppPublishingState{
		"notpublished": ManagedAndroidLobAppPublishingState_NotPublished,
		"processing":   ManagedAndroidLobAppPublishingState_Processing,
		"published":    ManagedAndroidLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppPublishingState(input)
	return &out, nil
}
