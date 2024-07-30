package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelAssignmentAssignmentMethod string

const (
	SensitivityLabelAssignmentAssignmentMethod_Auto       SensitivityLabelAssignmentAssignmentMethod = "auto"
	SensitivityLabelAssignmentAssignmentMethod_Privileged SensitivityLabelAssignmentAssignmentMethod = "privileged"
	SensitivityLabelAssignmentAssignmentMethod_Standard   SensitivityLabelAssignmentAssignmentMethod = "standard"
)

func PossibleValuesForSensitivityLabelAssignmentAssignmentMethod() []string {
	return []string{
		string(SensitivityLabelAssignmentAssignmentMethod_Auto),
		string(SensitivityLabelAssignmentAssignmentMethod_Privileged),
		string(SensitivityLabelAssignmentAssignmentMethod_Standard),
	}
}

func (s *SensitivityLabelAssignmentAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelAssignmentAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelAssignmentAssignmentMethod(input string) (*SensitivityLabelAssignmentAssignmentMethod, error) {
	vals := map[string]SensitivityLabelAssignmentAssignmentMethod{
		"auto":       SensitivityLabelAssignmentAssignmentMethod_Auto,
		"privileged": SensitivityLabelAssignmentAssignmentMethod_Privileged,
		"standard":   SensitivityLabelAssignmentAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelAssignmentAssignmentMethod(input)
	return &out, nil
}
