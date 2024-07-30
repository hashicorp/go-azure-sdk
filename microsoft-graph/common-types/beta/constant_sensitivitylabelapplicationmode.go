package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelApplicationMode string

const (
	SensitivityLabelApplicationMode_Automatic   SensitivityLabelApplicationMode = "automatic"
	SensitivityLabelApplicationMode_Manual      SensitivityLabelApplicationMode = "manual"
	SensitivityLabelApplicationMode_Recommended SensitivityLabelApplicationMode = "recommended"
)

func PossibleValuesForSensitivityLabelApplicationMode() []string {
	return []string{
		string(SensitivityLabelApplicationMode_Automatic),
		string(SensitivityLabelApplicationMode_Manual),
		string(SensitivityLabelApplicationMode_Recommended),
	}
}

func (s *SensitivityLabelApplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelApplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelApplicationMode(input string) (*SensitivityLabelApplicationMode, error) {
	vals := map[string]SensitivityLabelApplicationMode{
		"automatic":   SensitivityLabelApplicationMode_Automatic,
		"manual":      SensitivityLabelApplicationMode_Manual,
		"recommended": SensitivityLabelApplicationMode_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelApplicationMode(input)
	return &out, nil
}
