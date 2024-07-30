package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWebAppPublishingState string

const (
	WindowsWebAppPublishingState_NotPublished WindowsWebAppPublishingState = "notPublished"
	WindowsWebAppPublishingState_Processing   WindowsWebAppPublishingState = "processing"
	WindowsWebAppPublishingState_Published    WindowsWebAppPublishingState = "published"
)

func PossibleValuesForWindowsWebAppPublishingState() []string {
	return []string{
		string(WindowsWebAppPublishingState_NotPublished),
		string(WindowsWebAppPublishingState_Processing),
		string(WindowsWebAppPublishingState_Published),
	}
}

func (s *WindowsWebAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWebAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWebAppPublishingState(input string) (*WindowsWebAppPublishingState, error) {
	vals := map[string]WindowsWebAppPublishingState{
		"notpublished": WindowsWebAppPublishingState_NotPublished,
		"processing":   WindowsWebAppPublishingState_Processing,
		"published":    WindowsWebAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWebAppPublishingState(input)
	return &out, nil
}
