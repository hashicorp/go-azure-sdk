package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppPublishingState string

const (
	WebAppPublishingState_NotPublished WebAppPublishingState = "notPublished"
	WebAppPublishingState_Processing   WebAppPublishingState = "processing"
	WebAppPublishingState_Published    WebAppPublishingState = "published"
)

func PossibleValuesForWebAppPublishingState() []string {
	return []string{
		string(WebAppPublishingState_NotPublished),
		string(WebAppPublishingState_Processing),
		string(WebAppPublishingState_Published),
	}
}

func (s *WebAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebAppPublishingState(input string) (*WebAppPublishingState, error) {
	vals := map[string]WebAppPublishingState{
		"notpublished": WebAppPublishingState_NotPublished,
		"processing":   WebAppPublishingState_Processing,
		"published":    WebAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebAppPublishingState(input)
	return &out, nil
}
