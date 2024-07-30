package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWebClipPublishingState string

const (
	MacOSWebClipPublishingState_NotPublished MacOSWebClipPublishingState = "notPublished"
	MacOSWebClipPublishingState_Processing   MacOSWebClipPublishingState = "processing"
	MacOSWebClipPublishingState_Published    MacOSWebClipPublishingState = "published"
)

func PossibleValuesForMacOSWebClipPublishingState() []string {
	return []string{
		string(MacOSWebClipPublishingState_NotPublished),
		string(MacOSWebClipPublishingState_Processing),
		string(MacOSWebClipPublishingState_Published),
	}
}

func (s *MacOSWebClipPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWebClipPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWebClipPublishingState(input string) (*MacOSWebClipPublishingState, error) {
	vals := map[string]MacOSWebClipPublishingState{
		"notpublished": MacOSWebClipPublishingState_NotPublished,
		"processing":   MacOSWebClipPublishingState_Processing,
		"published":    MacOSWebClipPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWebClipPublishingState(input)
	return &out, nil
}
