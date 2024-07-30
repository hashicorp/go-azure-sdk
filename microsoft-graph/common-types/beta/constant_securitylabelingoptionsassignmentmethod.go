package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelingOptionsAssignmentMethod string

const (
	SecurityLabelingOptionsAssignmentMethod_Auto       SecurityLabelingOptionsAssignmentMethod = "auto"
	SecurityLabelingOptionsAssignmentMethod_Privileged SecurityLabelingOptionsAssignmentMethod = "privileged"
	SecurityLabelingOptionsAssignmentMethod_Standard   SecurityLabelingOptionsAssignmentMethod = "standard"
)

func PossibleValuesForSecurityLabelingOptionsAssignmentMethod() []string {
	return []string{
		string(SecurityLabelingOptionsAssignmentMethod_Auto),
		string(SecurityLabelingOptionsAssignmentMethod_Privileged),
		string(SecurityLabelingOptionsAssignmentMethod_Standard),
	}
}

func (s *SecurityLabelingOptionsAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityLabelingOptionsAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityLabelingOptionsAssignmentMethod(input string) (*SecurityLabelingOptionsAssignmentMethod, error) {
	vals := map[string]SecurityLabelingOptionsAssignmentMethod{
		"auto":       SecurityLabelingOptionsAssignmentMethod_Auto,
		"privileged": SecurityLabelingOptionsAssignmentMethod_Privileged,
		"standard":   SecurityLabelingOptionsAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityLabelingOptionsAssignmentMethod(input)
	return &out, nil
}
