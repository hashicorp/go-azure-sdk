package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoot struct {
	ActivityBasedTimeoutPolicies              *[]ActivityBasedTimeoutPolicy              `json:"activityBasedTimeoutPolicies,omitempty"`
	AdminConsentRequestPolicy                 *AdminConsentRequestPolicy                 `json:"adminConsentRequestPolicy,omitempty"`
	AppManagementPolicies                     *[]AppManagementPolicy                     `json:"appManagementPolicies,omitempty"`
	AuthenticationFlowsPolicy                 *AuthenticationFlowsPolicy                 `json:"authenticationFlowsPolicy,omitempty"`
	AuthenticationMethodsPolicy               *AuthenticationMethodsPolicy               `json:"authenticationMethodsPolicy,omitempty"`
	AuthenticationStrengthPolicies            *[]AuthenticationStrengthPolicy            `json:"authenticationStrengthPolicies,omitempty"`
	AuthorizationPolicy                       *AuthorizationPolicy                       `json:"authorizationPolicy,omitempty"`
	ClaimsMappingPolicies                     *[]ClaimsMappingPolicy                     `json:"claimsMappingPolicies,omitempty"`
	ConditionalAccessPolicies                 *[]ConditionalAccessPolicy                 `json:"conditionalAccessPolicies,omitempty"`
	CrossTenantAccessPolicy                   *CrossTenantAccessPolicy                   `json:"crossTenantAccessPolicy,omitempty"`
	DefaultAppManagementPolicy                *TenantAppManagementPolicy                 `json:"defaultAppManagementPolicy,omitempty"`
	FeatureRolloutPolicies                    *[]FeatureRolloutPolicy                    `json:"featureRolloutPolicies,omitempty"`
	HomeRealmDiscoveryPolicies                *[]HomeRealmDiscoveryPolicy                `json:"homeRealmDiscoveryPolicies,omitempty"`
	Id                                        *string                                    `json:"id,omitempty"`
	IdentitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`
	ODataType                                 *string                                    `json:"@odata.type,omitempty"`
	PermissionGrantPolicies                   *[]PermissionGrantPolicy                   `json:"permissionGrantPolicies,omitempty"`
	RoleManagementPolicies                    *[]UnifiedRoleManagementPolicy             `json:"roleManagementPolicies,omitempty"`
	RoleManagementPolicyAssignments           *[]UnifiedRoleManagementPolicyAssignment   `json:"roleManagementPolicyAssignments,omitempty"`
	TokenIssuancePolicies                     *[]TokenIssuancePolicy                     `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies                     *[]TokenLifetimePolicy                     `json:"tokenLifetimePolicies,omitempty"`
}
