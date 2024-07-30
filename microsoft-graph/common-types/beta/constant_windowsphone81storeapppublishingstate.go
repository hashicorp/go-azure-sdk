package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81StoreAppPublishingState string

const (
	WindowsPhone81StoreAppPublishingState_NotPublished WindowsPhone81StoreAppPublishingState = "notPublished"
	WindowsPhone81StoreAppPublishingState_Processing   WindowsPhone81StoreAppPublishingState = "processing"
	WindowsPhone81StoreAppPublishingState_Published    WindowsPhone81StoreAppPublishingState = "published"
)

func PossibleValuesForWindowsPhone81StoreAppPublishingState() []string {
	return []string{
		string(WindowsPhone81StoreAppPublishingState_NotPublished),
		string(WindowsPhone81StoreAppPublishingState_Processing),
		string(WindowsPhone81StoreAppPublishingState_Published),
	}
}

func (s *WindowsPhone81StoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81StoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81StoreAppPublishingState(input string) (*WindowsPhone81StoreAppPublishingState, error) {
	vals := map[string]WindowsPhone81StoreAppPublishingState{
		"notpublished": WindowsPhone81StoreAppPublishingState_NotPublished,
		"processing":   WindowsPhone81StoreAppPublishingState_Processing,
		"published":    WindowsPhone81StoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81StoreAppPublishingState(input)
	return &out, nil
}
