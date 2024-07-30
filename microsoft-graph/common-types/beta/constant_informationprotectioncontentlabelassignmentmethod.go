package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionContentLabelAssignmentMethod string

const (
	InformationProtectionContentLabelAssignmentMethod_Auto       InformationProtectionContentLabelAssignmentMethod = "auto"
	InformationProtectionContentLabelAssignmentMethod_Privileged InformationProtectionContentLabelAssignmentMethod = "privileged"
	InformationProtectionContentLabelAssignmentMethod_Standard   InformationProtectionContentLabelAssignmentMethod = "standard"
)

func PossibleValuesForInformationProtectionContentLabelAssignmentMethod() []string {
	return []string{
		string(InformationProtectionContentLabelAssignmentMethod_Auto),
		string(InformationProtectionContentLabelAssignmentMethod_Privileged),
		string(InformationProtectionContentLabelAssignmentMethod_Standard),
	}
}

func (s *InformationProtectionContentLabelAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInformationProtectionContentLabelAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInformationProtectionContentLabelAssignmentMethod(input string) (*InformationProtectionContentLabelAssignmentMethod, error) {
	vals := map[string]InformationProtectionContentLabelAssignmentMethod{
		"auto":       InformationProtectionContentLabelAssignmentMethod_Auto,
		"privileged": InformationProtectionContentLabelAssignmentMethod_Privileged,
		"standard":   InformationProtectionContentLabelAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InformationProtectionContentLabelAssignmentMethod(input)
	return &out, nil
}
