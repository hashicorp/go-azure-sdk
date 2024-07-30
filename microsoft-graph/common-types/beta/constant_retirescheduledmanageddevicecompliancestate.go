package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceComplianceState string

const (
	RetireScheduledManagedDeviceComplianceState_Compliant     RetireScheduledManagedDeviceComplianceState = "compliant"
	RetireScheduledManagedDeviceComplianceState_Conflict      RetireScheduledManagedDeviceComplianceState = "conflict"
	RetireScheduledManagedDeviceComplianceState_Error         RetireScheduledManagedDeviceComplianceState = "error"
	RetireScheduledManagedDeviceComplianceState_NonCompliant  RetireScheduledManagedDeviceComplianceState = "nonCompliant"
	RetireScheduledManagedDeviceComplianceState_NotApplicable RetireScheduledManagedDeviceComplianceState = "notApplicable"
	RetireScheduledManagedDeviceComplianceState_NotAssigned   RetireScheduledManagedDeviceComplianceState = "notAssigned"
	RetireScheduledManagedDeviceComplianceState_Remediated    RetireScheduledManagedDeviceComplianceState = "remediated"
	RetireScheduledManagedDeviceComplianceState_Unknown       RetireScheduledManagedDeviceComplianceState = "unknown"
)

func PossibleValuesForRetireScheduledManagedDeviceComplianceState() []string {
	return []string{
		string(RetireScheduledManagedDeviceComplianceState_Compliant),
		string(RetireScheduledManagedDeviceComplianceState_Conflict),
		string(RetireScheduledManagedDeviceComplianceState_Error),
		string(RetireScheduledManagedDeviceComplianceState_NonCompliant),
		string(RetireScheduledManagedDeviceComplianceState_NotApplicable),
		string(RetireScheduledManagedDeviceComplianceState_NotAssigned),
		string(RetireScheduledManagedDeviceComplianceState_Remediated),
		string(RetireScheduledManagedDeviceComplianceState_Unknown),
	}
}

func (s *RetireScheduledManagedDeviceComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetireScheduledManagedDeviceComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetireScheduledManagedDeviceComplianceState(input string) (*RetireScheduledManagedDeviceComplianceState, error) {
	vals := map[string]RetireScheduledManagedDeviceComplianceState{
		"compliant":     RetireScheduledManagedDeviceComplianceState_Compliant,
		"conflict":      RetireScheduledManagedDeviceComplianceState_Conflict,
		"error":         RetireScheduledManagedDeviceComplianceState_Error,
		"noncompliant":  RetireScheduledManagedDeviceComplianceState_NonCompliant,
		"notapplicable": RetireScheduledManagedDeviceComplianceState_NotApplicable,
		"notassigned":   RetireScheduledManagedDeviceComplianceState_NotAssigned,
		"remediated":    RetireScheduledManagedDeviceComplianceState_Remediated,
		"unknown":       RetireScheduledManagedDeviceComplianceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceComplianceState(input)
	return &out, nil
}
