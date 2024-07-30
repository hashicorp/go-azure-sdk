package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionStatus string

const (
	SecurityActionStatus_Completed  SecurityActionStatus = "Completed"
	SecurityActionStatus_Failed     SecurityActionStatus = "Failed"
	SecurityActionStatus_NotStarted SecurityActionStatus = "NotStarted"
	SecurityActionStatus_Running    SecurityActionStatus = "Running"
)

func PossibleValuesForSecurityActionStatus() []string {
	return []string{
		string(SecurityActionStatus_Completed),
		string(SecurityActionStatus_Failed),
		string(SecurityActionStatus_NotStarted),
		string(SecurityActionStatus_Running),
	}
}

func (s *SecurityActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityActionStatus(input string) (*SecurityActionStatus, error) {
	vals := map[string]SecurityActionStatus{
		"completed":  SecurityActionStatus_Completed,
		"failed":     SecurityActionStatus_Failed,
		"notstarted": SecurityActionStatus_NotStarted,
		"running":    SecurityActionStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionStatus(input)
	return &out, nil
}
