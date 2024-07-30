package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXPublishingState string

const (
	WindowsPhone81AppXPublishingState_NotPublished WindowsPhone81AppXPublishingState = "notPublished"
	WindowsPhone81AppXPublishingState_Processing   WindowsPhone81AppXPublishingState = "processing"
	WindowsPhone81AppXPublishingState_Published    WindowsPhone81AppXPublishingState = "published"
)

func PossibleValuesForWindowsPhone81AppXPublishingState() []string {
	return []string{
		string(WindowsPhone81AppXPublishingState_NotPublished),
		string(WindowsPhone81AppXPublishingState_Processing),
		string(WindowsPhone81AppXPublishingState_Published),
	}
}

func (s *WindowsPhone81AppXPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81AppXPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81AppXPublishingState(input string) (*WindowsPhone81AppXPublishingState, error) {
	vals := map[string]WindowsPhone81AppXPublishingState{
		"notpublished": WindowsPhone81AppXPublishingState_NotPublished,
		"processing":   WindowsPhone81AppXPublishingState_Processing,
		"published":    WindowsPhone81AppXPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXPublishingState(input)
	return &out, nil
}
