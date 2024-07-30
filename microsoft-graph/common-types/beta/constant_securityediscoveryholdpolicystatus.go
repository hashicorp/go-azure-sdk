package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldPolicyStatus string

const (
	SecurityEdiscoveryHoldPolicyStatus_Error   SecurityEdiscoveryHoldPolicyStatus = "error"
	SecurityEdiscoveryHoldPolicyStatus_Pending SecurityEdiscoveryHoldPolicyStatus = "pending"
	SecurityEdiscoveryHoldPolicyStatus_Success SecurityEdiscoveryHoldPolicyStatus = "success"
)

func PossibleValuesForSecurityEdiscoveryHoldPolicyStatus() []string {
	return []string{
		string(SecurityEdiscoveryHoldPolicyStatus_Error),
		string(SecurityEdiscoveryHoldPolicyStatus_Pending),
		string(SecurityEdiscoveryHoldPolicyStatus_Success),
	}
}

func (s *SecurityEdiscoveryHoldPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryHoldPolicyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryHoldPolicyStatus(input string) (*SecurityEdiscoveryHoldPolicyStatus, error) {
	vals := map[string]SecurityEdiscoveryHoldPolicyStatus{
		"error":   SecurityEdiscoveryHoldPolicyStatus_Error,
		"pending": SecurityEdiscoveryHoldPolicyStatus_Pending,
		"success": SecurityEdiscoveryHoldPolicyStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryHoldPolicyStatus(input)
	return &out, nil
}
