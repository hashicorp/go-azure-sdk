package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsStoreAppPublishingState string

const (
	WindowsStoreAppPublishingState_NotPublished WindowsStoreAppPublishingState = "notPublished"
	WindowsStoreAppPublishingState_Processing   WindowsStoreAppPublishingState = "processing"
	WindowsStoreAppPublishingState_Published    WindowsStoreAppPublishingState = "published"
)

func PossibleValuesForWindowsStoreAppPublishingState() []string {
	return []string{
		string(WindowsStoreAppPublishingState_NotPublished),
		string(WindowsStoreAppPublishingState_Processing),
		string(WindowsStoreAppPublishingState_Published),
	}
}

func (s *WindowsStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsStoreAppPublishingState(input string) (*WindowsStoreAppPublishingState, error) {
	vals := map[string]WindowsStoreAppPublishingState{
		"notpublished": WindowsStoreAppPublishingState_NotPublished,
		"processing":   WindowsStoreAppPublishingState_Processing,
		"published":    WindowsStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsStoreAppPublishingState(input)
	return &out, nil
}
