package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMobileMSIPublishingState string

const (
	WindowsMobileMSIPublishingState_NotPublished WindowsMobileMSIPublishingState = "notPublished"
	WindowsMobileMSIPublishingState_Processing   WindowsMobileMSIPublishingState = "processing"
	WindowsMobileMSIPublishingState_Published    WindowsMobileMSIPublishingState = "published"
)

func PossibleValuesForWindowsMobileMSIPublishingState() []string {
	return []string{
		string(WindowsMobileMSIPublishingState_NotPublished),
		string(WindowsMobileMSIPublishingState_Processing),
		string(WindowsMobileMSIPublishingState_Published),
	}
}

func (s *WindowsMobileMSIPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsMobileMSIPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsMobileMSIPublishingState(input string) (*WindowsMobileMSIPublishingState, error) {
	vals := map[string]WindowsMobileMSIPublishingState{
		"notpublished": WindowsMobileMSIPublishingState_NotPublished,
		"processing":   WindowsMobileMSIPublishingState_Processing,
		"published":    WindowsMobileMSIPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMobileMSIPublishingState(input)
	return &out, nil
}
