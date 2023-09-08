package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute string

const (
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributecreatedDateTime       IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "CreatedDateTime"
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeHireDate      IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "EmployeeHireDate"
	IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeLeaveDateTime IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute = "EmployeeLeaveDateTime"
)

func PossibleValuesForIdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute() []string {
	return []string{
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributecreatedDateTime),
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeHireDate),
		string(IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeLeaveDateTime),
	}
}

func parseIdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute(input string) (*IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute, error) {
	vals := map[string]IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute{
		"createddatetime":       IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributecreatedDateTime,
		"employeehiredate":      IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeHireDate,
		"employeeleavedatetime": IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttributeemployeeLeaveDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute(input)
	return &out, nil
}
