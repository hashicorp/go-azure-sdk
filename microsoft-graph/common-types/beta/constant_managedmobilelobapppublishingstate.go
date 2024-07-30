package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileLobAppPublishingState string

const (
	ManagedMobileLobAppPublishingState_NotPublished ManagedMobileLobAppPublishingState = "notPublished"
	ManagedMobileLobAppPublishingState_Processing   ManagedMobileLobAppPublishingState = "processing"
	ManagedMobileLobAppPublishingState_Published    ManagedMobileLobAppPublishingState = "published"
)

func PossibleValuesForManagedMobileLobAppPublishingState() []string {
	return []string{
		string(ManagedMobileLobAppPublishingState_NotPublished),
		string(ManagedMobileLobAppPublishingState_Processing),
		string(ManagedMobileLobAppPublishingState_Published),
	}
}

func (s *ManagedMobileLobAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedMobileLobAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedMobileLobAppPublishingState(input string) (*ManagedMobileLobAppPublishingState, error) {
	vals := map[string]ManagedMobileLobAppPublishingState{
		"notpublished": ManagedMobileLobAppPublishingState_NotPublished,
		"processing":   ManagedMobileLobAppPublishingState_Processing,
		"published":    ManagedMobileLobAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedMobileLobAppPublishingState(input)
	return &out, nil
}
