package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppPublishingState string

const (
	AndroidForWorkAppPublishingState_NotPublished AndroidForWorkAppPublishingState = "notPublished"
	AndroidForWorkAppPublishingState_Processing   AndroidForWorkAppPublishingState = "processing"
	AndroidForWorkAppPublishingState_Published    AndroidForWorkAppPublishingState = "published"
)

func PossibleValuesForAndroidForWorkAppPublishingState() []string {
	return []string{
		string(AndroidForWorkAppPublishingState_NotPublished),
		string(AndroidForWorkAppPublishingState_Processing),
		string(AndroidForWorkAppPublishingState_Published),
	}
}

func (s *AndroidForWorkAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkAppPublishingState(input string) (*AndroidForWorkAppPublishingState, error) {
	vals := map[string]AndroidForWorkAppPublishingState{
		"notpublished": AndroidForWorkAppPublishingState_NotPublished,
		"processing":   AndroidForWorkAppPublishingState_Processing,
		"published":    AndroidForWorkAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkAppPublishingState(input)
	return &out, nil
}
