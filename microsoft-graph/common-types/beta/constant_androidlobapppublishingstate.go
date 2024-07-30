package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidLobAppPublishingState string

const (
	AndroidLobAppPublishingState_NotPublished AndroidLobAppPublishingState = "notPublished"
	AndroidLobAppPublishingState_Processing   AndroidLobAppPublishingState = "processing"
	AndroidLobAppPublishingState_Published    AndroidLobAppPublishingState = "published"
)

func PossibleValuesForAndroidLobAppPublishingState() []string {
	return []string{
		string(AndroidLobAppPublishingState_NotPublished),
		string(AndroidLobAppPublishingState_Processing),
		string(AndroidLobAppPublishingState_Published),
	}
}

func (s *AndroidLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidLobAppPublishingState(input string) (*AndroidLobAppPublishingState, error) {
	vals := map[string]AndroidLobAppPublishingState{
		"notpublished": AndroidLobAppPublishingState_NotPublished,
		"processing":   AndroidLobAppPublishingState_Processing,
		"published":    AndroidLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidLobAppPublishingState(input)
	return &out, nil
}
