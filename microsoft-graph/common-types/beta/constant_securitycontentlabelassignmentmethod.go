package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentLabelAssignmentMethod string

const (
	SecurityContentLabelAssignmentMethod_Auto       SecurityContentLabelAssignmentMethod = "auto"
	SecurityContentLabelAssignmentMethod_Privileged SecurityContentLabelAssignmentMethod = "privileged"
	SecurityContentLabelAssignmentMethod_Standard   SecurityContentLabelAssignmentMethod = "standard"
)

func PossibleValuesForSecurityContentLabelAssignmentMethod() []string {
	return []string{
		string(SecurityContentLabelAssignmentMethod_Auto),
		string(SecurityContentLabelAssignmentMethod_Privileged),
		string(SecurityContentLabelAssignmentMethod_Standard),
	}
}

func (s *SecurityContentLabelAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContentLabelAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContentLabelAssignmentMethod(input string) (*SecurityContentLabelAssignmentMethod, error) {
	vals := map[string]SecurityContentLabelAssignmentMethod{
		"auto":       SecurityContentLabelAssignmentMethod_Auto,
		"privileged": SecurityContentLabelAssignmentMethod_Privileged,
		"standard":   SecurityContentLabelAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentLabelAssignmentMethod(input)
	return &out, nil
}
