package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXBundlePublishingState string

const (
	WindowsPhone81AppXBundlePublishingState_NotPublished WindowsPhone81AppXBundlePublishingState = "notPublished"
	WindowsPhone81AppXBundlePublishingState_Processing   WindowsPhone81AppXBundlePublishingState = "processing"
	WindowsPhone81AppXBundlePublishingState_Published    WindowsPhone81AppXBundlePublishingState = "published"
)

func PossibleValuesForWindowsPhone81AppXBundlePublishingState() []string {
	return []string{
		string(WindowsPhone81AppXBundlePublishingState_NotPublished),
		string(WindowsPhone81AppXBundlePublishingState_Processing),
		string(WindowsPhone81AppXBundlePublishingState_Published),
	}
}

func (s *WindowsPhone81AppXBundlePublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81AppXBundlePublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81AppXBundlePublishingState(input string) (*WindowsPhone81AppXBundlePublishingState, error) {
	vals := map[string]WindowsPhone81AppXBundlePublishingState{
		"notpublished": WindowsPhone81AppXBundlePublishingState_NotPublished,
		"processing":   WindowsPhone81AppXBundlePublishingState_Processing,
		"published":    WindowsPhone81AppXBundlePublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXBundlePublishingState(input)
	return &out, nil
}
