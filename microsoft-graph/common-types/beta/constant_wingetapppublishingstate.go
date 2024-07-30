package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppPublishingState string

const (
	WinGetAppPublishingState_NotPublished WinGetAppPublishingState = "notPublished"
	WinGetAppPublishingState_Processing   WinGetAppPublishingState = "processing"
	WinGetAppPublishingState_Published    WinGetAppPublishingState = "published"
)

func PossibleValuesForWinGetAppPublishingState() []string {
	return []string{
		string(WinGetAppPublishingState_NotPublished),
		string(WinGetAppPublishingState_Processing),
		string(WinGetAppPublishingState_Published),
	}
}

func (s *WinGetAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWinGetAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWinGetAppPublishingState(input string) (*WinGetAppPublishingState, error) {
	vals := map[string]WinGetAppPublishingState{
		"notpublished": WinGetAppPublishingState_NotPublished,
		"processing":   WinGetAppPublishingState_Processing,
		"published":    WinGetAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WinGetAppPublishingState(input)
	return &out, nil
}
