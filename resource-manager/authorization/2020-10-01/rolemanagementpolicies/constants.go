package rolemanagementpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyRuleType string

const (
	RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule              RoleManagementPolicyRuleType = "RoleManagementPolicyApprovalRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule RoleManagementPolicyRuleType = "RoleManagementPolicyAuthenticationContextRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule            RoleManagementPolicyRuleType = "RoleManagementPolicyEnablementRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule            RoleManagementPolicyRuleType = "RoleManagementPolicyExpirationRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule          RoleManagementPolicyRuleType = "RoleManagementPolicyNotificationRule"
)

func PossibleValuesForRoleManagementPolicyRuleType() []string {
	return []string{
		string(RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule),
		string(RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule),
		string(RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
		string(RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
		string(RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	}
}
