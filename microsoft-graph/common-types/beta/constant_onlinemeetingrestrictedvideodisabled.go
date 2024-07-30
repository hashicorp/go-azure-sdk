package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRestrictedVideoDisabled string

const (
	OnlineMeetingRestrictedVideoDisabled_WatermarkProtection OnlineMeetingRestrictedVideoDisabled = "watermarkProtection"
)

func PossibleValuesForOnlineMeetingRestrictedVideoDisabled() []string {
	return []string{
		string(OnlineMeetingRestrictedVideoDisabled_WatermarkProtection),
	}
}

func (s *OnlineMeetingRestrictedVideoDisabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingRestrictedVideoDisabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingRestrictedVideoDisabled(input string) (*OnlineMeetingRestrictedVideoDisabled, error) {
	vals := map[string]OnlineMeetingRestrictedVideoDisabled{
		"watermarkprotection": OnlineMeetingRestrictedVideoDisabled_WatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingRestrictedVideoDisabled(input)
	return &out, nil
}
