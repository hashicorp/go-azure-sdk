package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosiPadOSWebClipPublishingState string

const (
	IosiPadOSWebClipPublishingState_NotPublished IosiPadOSWebClipPublishingState = "notPublished"
	IosiPadOSWebClipPublishingState_Processing   IosiPadOSWebClipPublishingState = "processing"
	IosiPadOSWebClipPublishingState_Published    IosiPadOSWebClipPublishingState = "published"
)

func PossibleValuesForIosiPadOSWebClipPublishingState() []string {
	return []string{
		string(IosiPadOSWebClipPublishingState_NotPublished),
		string(IosiPadOSWebClipPublishingState_Processing),
		string(IosiPadOSWebClipPublishingState_Published),
	}
}

func (s *IosiPadOSWebClipPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosiPadOSWebClipPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosiPadOSWebClipPublishingState(input string) (*IosiPadOSWebClipPublishingState, error) {
	vals := map[string]IosiPadOSWebClipPublishingState{
		"notpublished": IosiPadOSWebClipPublishingState_NotPublished,
		"processing":   IosiPadOSWebClipPublishingState_Processing,
		"published":    IosiPadOSWebClipPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosiPadOSWebClipPublishingState(input)
	return &out, nil
}
