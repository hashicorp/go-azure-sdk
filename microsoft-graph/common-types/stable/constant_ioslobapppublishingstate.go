package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppPublishingState string

const (
	IosLobAppPublishingState_NotPublished IosLobAppPublishingState = "notPublished"
	IosLobAppPublishingState_Processing   IosLobAppPublishingState = "processing"
	IosLobAppPublishingState_Published    IosLobAppPublishingState = "published"
)

func PossibleValuesForIosLobAppPublishingState() []string {
	return []string{
		string(IosLobAppPublishingState_NotPublished),
		string(IosLobAppPublishingState_Processing),
		string(IosLobAppPublishingState_Published),
	}
}

func (s *IosLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosLobAppPublishingState(input string) (*IosLobAppPublishingState, error) {
	vals := map[string]IosLobAppPublishingState{
		"notpublished": IosLobAppPublishingState_NotPublished,
		"processing":   IosLobAppPublishingState_Processing,
		"published":    IosLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosLobAppPublishingState(input)
	return &out, nil
}
