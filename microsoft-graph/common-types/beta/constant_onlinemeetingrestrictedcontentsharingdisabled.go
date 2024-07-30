package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRestrictedContentSharingDisabled string

const (
	OnlineMeetingRestrictedContentSharingDisabled_WatermarkProtection OnlineMeetingRestrictedContentSharingDisabled = "watermarkProtection"
)

func PossibleValuesForOnlineMeetingRestrictedContentSharingDisabled() []string {
	return []string{
		string(OnlineMeetingRestrictedContentSharingDisabled_WatermarkProtection),
	}
}

func (s *OnlineMeetingRestrictedContentSharingDisabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingRestrictedContentSharingDisabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingRestrictedContentSharingDisabled(input string) (*OnlineMeetingRestrictedContentSharingDisabled, error) {
	vals := map[string]OnlineMeetingRestrictedContentSharingDisabled{
		"watermarkprotection": OnlineMeetingRestrictedContentSharingDisabled_WatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingRestrictedContentSharingDisabled(input)
	return &out, nil
}
