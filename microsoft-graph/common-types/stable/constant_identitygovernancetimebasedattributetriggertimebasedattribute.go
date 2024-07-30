package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute string

const (
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_CreatedDateTime       IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "createdDateTime"
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeHireDate      IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "employeeHireDate"
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeLeaveDateTime IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "employeeLeaveDateTime"
)

func PossibleValuesForIdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute() []string {
	return []string{
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_CreatedDateTime),
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeHireDate),
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeLeaveDateTime),
	}
}

func (s *IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute(input string) (*IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute, error) {
	vals := map[string]IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute{
		"createddatetime":       IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_CreatedDateTime,
		"employeehiredate":      IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeHireDate,
		"employeeleavedatetime": IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute_EmployeeLeaveDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute(input)
	return &out, nil
}
