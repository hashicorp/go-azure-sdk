package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetailsStatus string

const (
	SecurityRunDetailsStatus_Completed       SecurityRunDetailsStatus = "completed"
	SecurityRunDetailsStatus_Failed          SecurityRunDetailsStatus = "failed"
	SecurityRunDetailsStatus_PartiallyFailed SecurityRunDetailsStatus = "partiallyFailed"
	SecurityRunDetailsStatus_Running         SecurityRunDetailsStatus = "running"
)

func PossibleValuesForSecurityRunDetailsStatus() []string {
	return []string{
		string(SecurityRunDetailsStatus_Completed),
		string(SecurityRunDetailsStatus_Failed),
		string(SecurityRunDetailsStatus_PartiallyFailed),
		string(SecurityRunDetailsStatus_Running),
	}
}

func (s *SecurityRunDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRunDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRunDetailsStatus(input string) (*SecurityRunDetailsStatus, error) {
	vals := map[string]SecurityRunDetailsStatus{
		"completed":       SecurityRunDetailsStatus_Completed,
		"failed":          SecurityRunDetailsStatus_Failed,
		"partiallyfailed": SecurityRunDetailsStatus_PartiallyFailed,
		"running":         SecurityRunDetailsStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunDetailsStatus(input)
	return &out, nil
}
