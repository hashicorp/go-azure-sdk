package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelingOptionsAssignmentMethod string

const (
	LabelingOptionsAssignmentMethod_Auto       LabelingOptionsAssignmentMethod = "auto"
	LabelingOptionsAssignmentMethod_Privileged LabelingOptionsAssignmentMethod = "privileged"
	LabelingOptionsAssignmentMethod_Standard   LabelingOptionsAssignmentMethod = "standard"
)

func PossibleValuesForLabelingOptionsAssignmentMethod() []string {
	return []string{
		string(LabelingOptionsAssignmentMethod_Auto),
		string(LabelingOptionsAssignmentMethod_Privileged),
		string(LabelingOptionsAssignmentMethod_Standard),
	}
}

func (s *LabelingOptionsAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabelingOptionsAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabelingOptionsAssignmentMethod(input string) (*LabelingOptionsAssignmentMethod, error) {
	vals := map[string]LabelingOptionsAssignmentMethod{
		"auto":       LabelingOptionsAssignmentMethod_Auto,
		"privileged": LabelingOptionsAssignmentMethod_Privileged,
		"standard":   LabelingOptionsAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelingOptionsAssignmentMethod(input)
	return &out, nil
}
