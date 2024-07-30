package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXPublishingState string

const (
	WindowsUniversalAppXPublishingState_NotPublished WindowsUniversalAppXPublishingState = "notPublished"
	WindowsUniversalAppXPublishingState_Processing   WindowsUniversalAppXPublishingState = "processing"
	WindowsUniversalAppXPublishingState_Published    WindowsUniversalAppXPublishingState = "published"
)

func PossibleValuesForWindowsUniversalAppXPublishingState() []string {
	return []string{
		string(WindowsUniversalAppXPublishingState_NotPublished),
		string(WindowsUniversalAppXPublishingState_Processing),
		string(WindowsUniversalAppXPublishingState_Published),
	}
}

func (s *WindowsUniversalAppXPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUniversalAppXPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUniversalAppXPublishingState(input string) (*WindowsUniversalAppXPublishingState, error) {
	vals := map[string]WindowsUniversalAppXPublishingState{
		"notpublished": WindowsUniversalAppXPublishingState_NotPublished,
		"processing":   WindowsUniversalAppXPublishingState_Processing,
		"published":    WindowsUniversalAppXPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXPublishingState(input)
	return &out, nil
}
