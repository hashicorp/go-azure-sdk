package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPublishingState string

const (
	Win32LobAppPublishingState_NotPublished Win32LobAppPublishingState = "notPublished"
	Win32LobAppPublishingState_Processing   Win32LobAppPublishingState = "processing"
	Win32LobAppPublishingState_Published    Win32LobAppPublishingState = "published"
)

func PossibleValuesForWin32LobAppPublishingState() []string {
	return []string{
		string(Win32LobAppPublishingState_NotPublished),
		string(Win32LobAppPublishingState_Processing),
		string(Win32LobAppPublishingState_Published),
	}
}

func (s *Win32LobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPublishingState(input string) (*Win32LobAppPublishingState, error) {
	vals := map[string]Win32LobAppPublishingState{
		"notpublished": Win32LobAppPublishingState_NotPublished,
		"processing":   Win32LobAppPublishingState_Processing,
		"published":    Win32LobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPublishingState(input)
	return &out, nil
}
