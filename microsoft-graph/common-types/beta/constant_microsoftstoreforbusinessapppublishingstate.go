package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessAppPublishingState string

const (
	MicrosoftStoreForBusinessAppPublishingState_NotPublished MicrosoftStoreForBusinessAppPublishingState = "notPublished"
	MicrosoftStoreForBusinessAppPublishingState_Processing   MicrosoftStoreForBusinessAppPublishingState = "processing"
	MicrosoftStoreForBusinessAppPublishingState_Published    MicrosoftStoreForBusinessAppPublishingState = "published"
)

func PossibleValuesForMicrosoftStoreForBusinessAppPublishingState() []string {
	return []string{
		string(MicrosoftStoreForBusinessAppPublishingState_NotPublished),
		string(MicrosoftStoreForBusinessAppPublishingState_Processing),
		string(MicrosoftStoreForBusinessAppPublishingState_Published),
	}
}

func (s *MicrosoftStoreForBusinessAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftStoreForBusinessAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftStoreForBusinessAppPublishingState(input string) (*MicrosoftStoreForBusinessAppPublishingState, error) {
	vals := map[string]MicrosoftStoreForBusinessAppPublishingState{
		"notpublished": MicrosoftStoreForBusinessAppPublishingState_NotPublished,
		"processing":   MicrosoftStoreForBusinessAppPublishingState_Processing,
		"published":    MicrosoftStoreForBusinessAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftStoreForBusinessAppPublishingState(input)
	return &out, nil
}
