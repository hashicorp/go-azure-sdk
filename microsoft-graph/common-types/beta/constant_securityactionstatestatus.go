package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionStateStatus string

const (
	SecurityActionStateStatus_Completed  SecurityActionStateStatus = "Completed"
	SecurityActionStateStatus_Failed     SecurityActionStateStatus = "Failed"
	SecurityActionStateStatus_NotStarted SecurityActionStateStatus = "NotStarted"
	SecurityActionStateStatus_Running    SecurityActionStateStatus = "Running"
)

func PossibleValuesForSecurityActionStateStatus() []string {
	return []string{
		string(SecurityActionStateStatus_Completed),
		string(SecurityActionStateStatus_Failed),
		string(SecurityActionStateStatus_NotStarted),
		string(SecurityActionStateStatus_Running),
	}
}

func (s *SecurityActionStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityActionStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityActionStateStatus(input string) (*SecurityActionStateStatus, error) {
	vals := map[string]SecurityActionStateStatus{
		"completed":  SecurityActionStateStatus_Completed,
		"failed":     SecurityActionStateStatus_Failed,
		"notstarted": SecurityActionStateStatus_NotStarted,
		"running":    SecurityActionStateStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionStateStatus(input)
	return &out, nil
}
