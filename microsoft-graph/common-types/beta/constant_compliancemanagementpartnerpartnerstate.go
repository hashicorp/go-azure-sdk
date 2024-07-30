package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartnerPartnerState string

const (
	ComplianceManagementPartnerPartnerState_Enabled      ComplianceManagementPartnerPartnerState = "enabled"
	ComplianceManagementPartnerPartnerState_Rejected     ComplianceManagementPartnerPartnerState = "rejected"
	ComplianceManagementPartnerPartnerState_Terminated   ComplianceManagementPartnerPartnerState = "terminated"
	ComplianceManagementPartnerPartnerState_Unavailable  ComplianceManagementPartnerPartnerState = "unavailable"
	ComplianceManagementPartnerPartnerState_Unknown      ComplianceManagementPartnerPartnerState = "unknown"
	ComplianceManagementPartnerPartnerState_Unresponsive ComplianceManagementPartnerPartnerState = "unresponsive"
)

func PossibleValuesForComplianceManagementPartnerPartnerState() []string {
	return []string{
		string(ComplianceManagementPartnerPartnerState_Enabled),
		string(ComplianceManagementPartnerPartnerState_Rejected),
		string(ComplianceManagementPartnerPartnerState_Terminated),
		string(ComplianceManagementPartnerPartnerState_Unavailable),
		string(ComplianceManagementPartnerPartnerState_Unknown),
		string(ComplianceManagementPartnerPartnerState_Unresponsive),
	}
}

func (s *ComplianceManagementPartnerPartnerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceManagementPartnerPartnerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceManagementPartnerPartnerState(input string) (*ComplianceManagementPartnerPartnerState, error) {
	vals := map[string]ComplianceManagementPartnerPartnerState{
		"enabled":      ComplianceManagementPartnerPartnerState_Enabled,
		"rejected":     ComplianceManagementPartnerPartnerState_Rejected,
		"terminated":   ComplianceManagementPartnerPartnerState_Terminated,
		"unavailable":  ComplianceManagementPartnerPartnerState_Unavailable,
		"unknown":      ComplianceManagementPartnerPartnerState_Unknown,
		"unresponsive": ComplianceManagementPartnerPartnerState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceManagementPartnerPartnerState(input)
	return &out, nil
}
