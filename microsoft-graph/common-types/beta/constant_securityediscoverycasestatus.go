package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCaseStatus string

const (
	SecurityEdiscoveryCaseStatus_Active          SecurityEdiscoveryCaseStatus = "active"
	SecurityEdiscoveryCaseStatus_Closed          SecurityEdiscoveryCaseStatus = "closed"
	SecurityEdiscoveryCaseStatus_ClosedWithError SecurityEdiscoveryCaseStatus = "closedWithError"
	SecurityEdiscoveryCaseStatus_Closing         SecurityEdiscoveryCaseStatus = "closing"
	SecurityEdiscoveryCaseStatus_PendingDelete   SecurityEdiscoveryCaseStatus = "pendingDelete"
	SecurityEdiscoveryCaseStatus_Unknown         SecurityEdiscoveryCaseStatus = "unknown"
)

func PossibleValuesForSecurityEdiscoveryCaseStatus() []string {
	return []string{
		string(SecurityEdiscoveryCaseStatus_Active),
		string(SecurityEdiscoveryCaseStatus_Closed),
		string(SecurityEdiscoveryCaseStatus_ClosedWithError),
		string(SecurityEdiscoveryCaseStatus_Closing),
		string(SecurityEdiscoveryCaseStatus_PendingDelete),
		string(SecurityEdiscoveryCaseStatus_Unknown),
	}
}

func (s *SecurityEdiscoveryCaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryCaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryCaseStatus(input string) (*SecurityEdiscoveryCaseStatus, error) {
	vals := map[string]SecurityEdiscoveryCaseStatus{
		"active":          SecurityEdiscoveryCaseStatus_Active,
		"closed":          SecurityEdiscoveryCaseStatus_Closed,
		"closedwitherror": SecurityEdiscoveryCaseStatus_ClosedWithError,
		"closing":         SecurityEdiscoveryCaseStatus_Closing,
		"pendingdelete":   SecurityEdiscoveryCaseStatus_PendingDelete,
		"unknown":         SecurityEdiscoveryCaseStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCaseStatus(input)
	return &out, nil
}
