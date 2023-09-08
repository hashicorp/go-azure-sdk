package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyRuleTargetOperations string

const (
	UnifiedRoleManagementPolicyRuleTargetOperationsactivate   UnifiedRoleManagementPolicyRuleTargetOperations = "Activate"
	UnifiedRoleManagementPolicyRuleTargetOperationsall        UnifiedRoleManagementPolicyRuleTargetOperations = "All"
	UnifiedRoleManagementPolicyRuleTargetOperationsassign     UnifiedRoleManagementPolicyRuleTargetOperations = "Assign"
	UnifiedRoleManagementPolicyRuleTargetOperationsdeactivate UnifiedRoleManagementPolicyRuleTargetOperations = "Deactivate"
	UnifiedRoleManagementPolicyRuleTargetOperationsextend     UnifiedRoleManagementPolicyRuleTargetOperations = "Extend"
	UnifiedRoleManagementPolicyRuleTargetOperationsremove     UnifiedRoleManagementPolicyRuleTargetOperations = "Remove"
	UnifiedRoleManagementPolicyRuleTargetOperationsrenew      UnifiedRoleManagementPolicyRuleTargetOperations = "Renew"
	UnifiedRoleManagementPolicyRuleTargetOperationsupdate     UnifiedRoleManagementPolicyRuleTargetOperations = "Update"
)

func PossibleValuesForUnifiedRoleManagementPolicyRuleTargetOperations() []string {
	return []string{
		string(UnifiedRoleManagementPolicyRuleTargetOperationsactivate),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsall),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsassign),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsdeactivate),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsextend),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsremove),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsrenew),
		string(UnifiedRoleManagementPolicyRuleTargetOperationsupdate),
	}
}

func parseUnifiedRoleManagementPolicyRuleTargetOperations(input string) (*UnifiedRoleManagementPolicyRuleTargetOperations, error) {
	vals := map[string]UnifiedRoleManagementPolicyRuleTargetOperations{
		"activate":   UnifiedRoleManagementPolicyRuleTargetOperationsactivate,
		"all":        UnifiedRoleManagementPolicyRuleTargetOperationsall,
		"assign":     UnifiedRoleManagementPolicyRuleTargetOperationsassign,
		"deactivate": UnifiedRoleManagementPolicyRuleTargetOperationsdeactivate,
		"extend":     UnifiedRoleManagementPolicyRuleTargetOperationsextend,
		"remove":     UnifiedRoleManagementPolicyRuleTargetOperationsremove,
		"renew":      UnifiedRoleManagementPolicyRuleTargetOperationsrenew,
		"update":     UnifiedRoleManagementPolicyRuleTargetOperationsupdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleManagementPolicyRuleTargetOperations(input)
	return &out, nil
}
