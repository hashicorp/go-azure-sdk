package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoot struct {
	AccessReviewPolicy                        *AccessReviewPolicy                        `json:"accessReviewPolicy,omitempty"`
	ActivityBasedTimeoutPolicies              *[]ActivityBasedTimeoutPolicy              `json:"activityBasedTimeoutPolicies,omitempty"`
	AdminConsentRequestPolicy                 *AdminConsentRequestPolicy                 `json:"adminConsentRequestPolicy,omitempty"`
	AppManagementPolicies                     *[]AppManagementPolicy                     `json:"appManagementPolicies,omitempty"`
	AuthenticationFlowsPolicy                 *AuthenticationFlowsPolicy                 `json:"authenticationFlowsPolicy,omitempty"`
	AuthenticationMethodsPolicy               *AuthenticationMethodsPolicy               `json:"authenticationMethodsPolicy,omitempty"`
	AuthenticationStrengthPolicies            *[]AuthenticationStrengthPolicy            `json:"authenticationStrengthPolicies,omitempty"`
	AuthorizationPolicy                       *[]AuthorizationPolicy                     `json:"authorizationPolicy,omitempty"`
	B2cAuthenticationMethodsPolicy            *B2cAuthenticationMethodsPolicy            `json:"b2cAuthenticationMethodsPolicy,omitempty"`
	ClaimsMappingPolicies                     *[]ClaimsMappingPolicy                     `json:"claimsMappingPolicies,omitempty"`
	ConditionalAccessPolicies                 *[]ConditionalAccessPolicy                 `json:"conditionalAccessPolicies,omitempty"`
	CrossTenantAccessPolicy                   *CrossTenantAccessPolicy                   `json:"crossTenantAccessPolicy,omitempty"`
	DefaultAppManagementPolicy                *TenantAppManagementPolicy                 `json:"defaultAppManagementPolicy,omitempty"`
	DeviceRegistrationPolicy                  *DeviceRegistrationPolicy                  `json:"deviceRegistrationPolicy,omitempty"`
	DirectoryRoleAccessReviewPolicy           *DirectoryRoleAccessReviewPolicy           `json:"directoryRoleAccessReviewPolicy,omitempty"`
	ExternalIdentitiesPolicy                  *ExternalIdentitiesPolicy                  `json:"externalIdentitiesPolicy,omitempty"`
	FeatureRolloutPolicies                    *[]FeatureRolloutPolicy                    `json:"featureRolloutPolicies,omitempty"`
	FederatedTokenValidationPolicy            *FederatedTokenValidationPolicy            `json:"federatedTokenValidationPolicy,omitempty"`
	HomeRealmDiscoveryPolicies                *[]HomeRealmDiscoveryPolicy                `json:"homeRealmDiscoveryPolicies,omitempty"`
	IdentitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`
	MobileAppManagementPolicies               *[]MobilityManagementPolicy                `json:"mobileAppManagementPolicies,omitempty"`
	MobileDeviceManagementPolicies            *[]MobilityManagementPolicy                `json:"mobileDeviceManagementPolicies,omitempty"`
	ODataType                                 *string                                    `json:"@odata.type,omitempty"`
	PermissionGrantPolicies                   *[]PermissionGrantPolicy                   `json:"permissionGrantPolicies,omitempty"`
	RoleManagementPolicies                    *[]UnifiedRoleManagementPolicy             `json:"roleManagementPolicies,omitempty"`
	RoleManagementPolicyAssignments           *[]UnifiedRoleManagementPolicyAssignment   `json:"roleManagementPolicyAssignments,omitempty"`
	ServicePrincipalCreationPolicies          *[]ServicePrincipalCreationPolicy          `json:"servicePrincipalCreationPolicies,omitempty"`
	TokenIssuancePolicies                     *[]TokenIssuancePolicy                     `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies                     *[]TokenLifetimePolicy                     `json:"tokenLifetimePolicies,omitempty"`
}
