package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartnerPartnerState string

const (
	ComplianceManagementPartnerPartnerStateenabled      ComplianceManagementPartnerPartnerState = "Enabled"
	ComplianceManagementPartnerPartnerStaterejected     ComplianceManagementPartnerPartnerState = "Rejected"
	ComplianceManagementPartnerPartnerStateterminated   ComplianceManagementPartnerPartnerState = "Terminated"
	ComplianceManagementPartnerPartnerStateunavailable  ComplianceManagementPartnerPartnerState = "Unavailable"
	ComplianceManagementPartnerPartnerStateunknown      ComplianceManagementPartnerPartnerState = "Unknown"
	ComplianceManagementPartnerPartnerStateunresponsive ComplianceManagementPartnerPartnerState = "Unresponsive"
)

func PossibleValuesForComplianceManagementPartnerPartnerState() []string {
	return []string{
		string(ComplianceManagementPartnerPartnerStateenabled),
		string(ComplianceManagementPartnerPartnerStaterejected),
		string(ComplianceManagementPartnerPartnerStateterminated),
		string(ComplianceManagementPartnerPartnerStateunavailable),
		string(ComplianceManagementPartnerPartnerStateunknown),
		string(ComplianceManagementPartnerPartnerStateunresponsive),
	}
}

func parseComplianceManagementPartnerPartnerState(input string) (*ComplianceManagementPartnerPartnerState, error) {
	vals := map[string]ComplianceManagementPartnerPartnerState{
		"enabled":      ComplianceManagementPartnerPartnerStateenabled,
		"rejected":     ComplianceManagementPartnerPartnerStaterejected,
		"terminated":   ComplianceManagementPartnerPartnerStateterminated,
		"unavailable":  ComplianceManagementPartnerPartnerStateunavailable,
		"unknown":      ComplianceManagementPartnerPartnerStateunknown,
		"unresponsive": ComplianceManagementPartnerPartnerStateunresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceManagementPartnerPartnerState(input)
	return &out, nil
}
