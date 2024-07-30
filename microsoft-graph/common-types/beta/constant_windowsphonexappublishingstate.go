package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneXAPPublishingState string

const (
	WindowsPhoneXAPPublishingState_NotPublished WindowsPhoneXAPPublishingState = "notPublished"
	WindowsPhoneXAPPublishingState_Processing   WindowsPhoneXAPPublishingState = "processing"
	WindowsPhoneXAPPublishingState_Published    WindowsPhoneXAPPublishingState = "published"
)

func PossibleValuesForWindowsPhoneXAPPublishingState() []string {
	return []string{
		string(WindowsPhoneXAPPublishingState_NotPublished),
		string(WindowsPhoneXAPPublishingState_Processing),
		string(WindowsPhoneXAPPublishingState_Published),
	}
}

func (s *WindowsPhoneXAPPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneXAPPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneXAPPublishingState(input string) (*WindowsPhoneXAPPublishingState, error) {
	vals := map[string]WindowsPhoneXAPPublishingState{
		"notpublished": WindowsPhoneXAPPublishingState_NotPublished,
		"processing":   WindowsPhoneXAPPublishingState_Processing,
		"published":    WindowsPhoneXAPPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneXAPPublishingState(input)
	return &out, nil
}
