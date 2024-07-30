package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppXPublishingState string

const (
	WindowsAppXPublishingState_NotPublished WindowsAppXPublishingState = "notPublished"
	WindowsAppXPublishingState_Processing   WindowsAppXPublishingState = "processing"
	WindowsAppXPublishingState_Published    WindowsAppXPublishingState = "published"
)

func PossibleValuesForWindowsAppXPublishingState() []string {
	return []string{
		string(WindowsAppXPublishingState_NotPublished),
		string(WindowsAppXPublishingState_Processing),
		string(WindowsAppXPublishingState_Published),
	}
}

func (s *WindowsAppXPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAppXPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAppXPublishingState(input string) (*WindowsAppXPublishingState, error) {
	vals := map[string]WindowsAppXPublishingState{
		"notpublished": WindowsAppXPublishingState_NotPublished,
		"processing":   WindowsAppXPublishingState_Processing,
		"published":    WindowsAppXPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAppXPublishingState(input)
	return &out, nil
}
