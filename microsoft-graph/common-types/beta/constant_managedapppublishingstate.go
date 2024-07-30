package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPublishingState string

const (
	ManagedAppPublishingState_NotPublished ManagedAppPublishingState = "notPublished"
	ManagedAppPublishingState_Processing   ManagedAppPublishingState = "processing"
	ManagedAppPublishingState_Published    ManagedAppPublishingState = "published"
)

func PossibleValuesForManagedAppPublishingState() []string {
	return []string{
		string(ManagedAppPublishingState_NotPublished),
		string(ManagedAppPublishingState_Processing),
		string(ManagedAppPublishingState_Published),
	}
}

func (s *ManagedAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppPublishingState(input string) (*ManagedAppPublishingState, error) {
	vals := map[string]ManagedAppPublishingState{
		"notpublished": ManagedAppPublishingState_NotPublished,
		"processing":   ManagedAppPublishingState_Processing,
		"published":    ManagedAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppPublishingState(input)
	return &out, nil
}
