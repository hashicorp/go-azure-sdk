package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMicrosoftEdgeAppPublishingState string

const (
	WindowsMicrosoftEdgeAppPublishingState_NotPublished WindowsMicrosoftEdgeAppPublishingState = "notPublished"
	WindowsMicrosoftEdgeAppPublishingState_Processing   WindowsMicrosoftEdgeAppPublishingState = "processing"
	WindowsMicrosoftEdgeAppPublishingState_Published    WindowsMicrosoftEdgeAppPublishingState = "published"
)

func PossibleValuesForWindowsMicrosoftEdgeAppPublishingState() []string {
	return []string{
		string(WindowsMicrosoftEdgeAppPublishingState_NotPublished),
		string(WindowsMicrosoftEdgeAppPublishingState_Processing),
		string(WindowsMicrosoftEdgeAppPublishingState_Published),
	}
}

func (s *WindowsMicrosoftEdgeAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsMicrosoftEdgeAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsMicrosoftEdgeAppPublishingState(input string) (*WindowsMicrosoftEdgeAppPublishingState, error) {
	vals := map[string]WindowsMicrosoftEdgeAppPublishingState{
		"notpublished": WindowsMicrosoftEdgeAppPublishingState_NotPublished,
		"processing":   WindowsMicrosoftEdgeAppPublishingState_Processing,
		"published":    WindowsMicrosoftEdgeAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMicrosoftEdgeAppPublishingState(input)
	return &out, nil
}
