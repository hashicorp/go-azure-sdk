package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPolicyBaseStatus string

const (
	SecurityPolicyBaseStatus_Error   SecurityPolicyBaseStatus = "error"
	SecurityPolicyBaseStatus_Pending SecurityPolicyBaseStatus = "pending"
	SecurityPolicyBaseStatus_Success SecurityPolicyBaseStatus = "success"
)

func PossibleValuesForSecurityPolicyBaseStatus() []string {
	return []string{
		string(SecurityPolicyBaseStatus_Error),
		string(SecurityPolicyBaseStatus_Pending),
		string(SecurityPolicyBaseStatus_Success),
	}
}

func (s *SecurityPolicyBaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityPolicyBaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityPolicyBaseStatus(input string) (*SecurityPolicyBaseStatus, error) {
	vals := map[string]SecurityPolicyBaseStatus{
		"error":   SecurityPolicyBaseStatus_Error,
		"pending": SecurityPolicyBaseStatus_Pending,
		"success": SecurityPolicyBaseStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityPolicyBaseStatus(input)
	return &out, nil
}
