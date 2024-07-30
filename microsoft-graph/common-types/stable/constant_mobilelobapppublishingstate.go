package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileLobAppPublishingState string

const (
	MobileLobAppPublishingState_NotPublished MobileLobAppPublishingState = "notPublished"
	MobileLobAppPublishingState_Processing   MobileLobAppPublishingState = "processing"
	MobileLobAppPublishingState_Published    MobileLobAppPublishingState = "published"
)

func PossibleValuesForMobileLobAppPublishingState() []string {
	return []string{
		string(MobileLobAppPublishingState_NotPublished),
		string(MobileLobAppPublishingState_Processing),
		string(MobileLobAppPublishingState_Published),
	}
}

func (s *MobileLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileLobAppPublishingState(input string) (*MobileLobAppPublishingState, error) {
	vals := map[string]MobileLobAppPublishingState{
		"notpublished": MobileLobAppPublishingState_NotPublished,
		"processing":   MobileLobAppPublishingState_Processing,
		"published":    MobileLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileLobAppPublishingState(input)
	return &out, nil
}
