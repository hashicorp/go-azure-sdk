package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationRequirementPolicyRequirementProvider string

const (
	AuthenticationRequirementPolicyRequirementProvideraccountCompromisePolicies                         AuthenticationRequirementPolicyRequirementProvider = "AccountCompromisePolicies"
	AuthenticationRequirementPolicyRequirementProviderbaselineProtection                                AuthenticationRequirementPolicyRequirementProvider = "BaselineProtection"
	AuthenticationRequirementPolicyRequirementProvidercrossTenantOutboundRule                           AuthenticationRequirementPolicyRequirementProvider = "CrossTenantOutboundRule"
	AuthenticationRequirementPolicyRequirementProviderenforcedForCspAdmins                              AuthenticationRequirementPolicyRequirementProvider = "EnforcedForCspAdmins"
	AuthenticationRequirementPolicyRequirementProvidergpsLocationCondition                              AuthenticationRequirementPolicyRequirementProvider = "GpsLocationCondition"
	AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByBaselineProtection       AuthenticationRequirementPolicyRequirementProvider = "MfaRegistrationRequiredByBaselineProtection"
	AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByIdentityProtectionPolicy AuthenticationRequirementPolicyRequirementProvider = "MfaRegistrationRequiredByIdentityProtectionPolicy"
	AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByMultiConditionalAccess   AuthenticationRequirementPolicyRequirementProvider = "MfaRegistrationRequiredByMultiConditionalAccess"
	AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredBySecurityDefaults         AuthenticationRequirementPolicyRequirementProvider = "MfaRegistrationRequiredBySecurityDefaults"
	AuthenticationRequirementPolicyRequirementProvidermultiConditionalAccess                            AuthenticationRequirementPolicyRequirementProvider = "MultiConditionalAccess"
	AuthenticationRequirementPolicyRequirementProviderproofUpCodeRequest                                AuthenticationRequirementPolicyRequirementProvider = "ProofUpCodeRequest"
	AuthenticationRequirementPolicyRequirementProviderrequest                                           AuthenticationRequirementPolicyRequirementProvider = "Request"
	AuthenticationRequirementPolicyRequirementProviderriskBasedPolicy                                   AuthenticationRequirementPolicyRequirementProvider = "RiskBasedPolicy"
	AuthenticationRequirementPolicyRequirementProvidersecurityDefaults                                  AuthenticationRequirementPolicyRequirementProvider = "SecurityDefaults"
	AuthenticationRequirementPolicyRequirementProviderservicePrincipal                                  AuthenticationRequirementPolicyRequirementProvider = "ServicePrincipal"
	AuthenticationRequirementPolicyRequirementProvidertenantSessionRiskPolicy                           AuthenticationRequirementPolicyRequirementProvider = "TenantSessionRiskPolicy"
	AuthenticationRequirementPolicyRequirementProvideruser                                              AuthenticationRequirementPolicyRequirementProvider = "User"
	AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccess                               AuthenticationRequirementPolicyRequirementProvider = "V1ConditionalAccess"
	AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessDependency                     AuthenticationRequirementPolicyRequirementProvider = "V1ConditionalAccessDependency"
	AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessPolicyIdRequested              AuthenticationRequirementPolicyRequirementProvider = "V1ConditionalAccessPolicyIdRequested"
)

func PossibleValuesForAuthenticationRequirementPolicyRequirementProvider() []string {
	return []string{
		string(AuthenticationRequirementPolicyRequirementProvideraccountCompromisePolicies),
		string(AuthenticationRequirementPolicyRequirementProviderbaselineProtection),
		string(AuthenticationRequirementPolicyRequirementProvidercrossTenantOutboundRule),
		string(AuthenticationRequirementPolicyRequirementProviderenforcedForCspAdmins),
		string(AuthenticationRequirementPolicyRequirementProvidergpsLocationCondition),
		string(AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByBaselineProtection),
		string(AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByIdentityProtectionPolicy),
		string(AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByMultiConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredBySecurityDefaults),
		string(AuthenticationRequirementPolicyRequirementProvidermultiConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProviderproofUpCodeRequest),
		string(AuthenticationRequirementPolicyRequirementProviderrequest),
		string(AuthenticationRequirementPolicyRequirementProviderriskBasedPolicy),
		string(AuthenticationRequirementPolicyRequirementProvidersecurityDefaults),
		string(AuthenticationRequirementPolicyRequirementProviderservicePrincipal),
		string(AuthenticationRequirementPolicyRequirementProvidertenantSessionRiskPolicy),
		string(AuthenticationRequirementPolicyRequirementProvideruser),
		string(AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessDependency),
		string(AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessPolicyIdRequested),
	}
}

func parseAuthenticationRequirementPolicyRequirementProvider(input string) (*AuthenticationRequirementPolicyRequirementProvider, error) {
	vals := map[string]AuthenticationRequirementPolicyRequirementProvider{
		"accountcompromisepolicies":                         AuthenticationRequirementPolicyRequirementProvideraccountCompromisePolicies,
		"baselineprotection":                                AuthenticationRequirementPolicyRequirementProviderbaselineProtection,
		"crosstenantoutboundrule":                           AuthenticationRequirementPolicyRequirementProvidercrossTenantOutboundRule,
		"enforcedforcspadmins":                              AuthenticationRequirementPolicyRequirementProviderenforcedForCspAdmins,
		"gpslocationcondition":                              AuthenticationRequirementPolicyRequirementProvidergpsLocationCondition,
		"mfaregistrationrequiredbybaselineprotection":       AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByBaselineProtection,
		"mfaregistrationrequiredbyidentityprotectionpolicy": AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByIdentityProtectionPolicy,
		"mfaregistrationrequiredbymulticonditionalaccess":   AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredByMultiConditionalAccess,
		"mfaregistrationrequiredbysecuritydefaults":         AuthenticationRequirementPolicyRequirementProvidermfaRegistrationRequiredBySecurityDefaults,
		"multiconditionalaccess":                            AuthenticationRequirementPolicyRequirementProvidermultiConditionalAccess,
		"proofupcoderequest":                                AuthenticationRequirementPolicyRequirementProviderproofUpCodeRequest,
		"request":                                           AuthenticationRequirementPolicyRequirementProviderrequest,
		"riskbasedpolicy":                                   AuthenticationRequirementPolicyRequirementProviderriskBasedPolicy,
		"securitydefaults":                                  AuthenticationRequirementPolicyRequirementProvidersecurityDefaults,
		"serviceprincipal":                                  AuthenticationRequirementPolicyRequirementProviderservicePrincipal,
		"tenantsessionriskpolicy":                           AuthenticationRequirementPolicyRequirementProvidertenantSessionRiskPolicy,
		"user":                                              AuthenticationRequirementPolicyRequirementProvideruser,
		"v1conditionalaccess":                               AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccess,
		"v1conditionalaccessdependency":                     AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessDependency,
		"v1conditionalaccesspolicyidrequested":              AuthenticationRequirementPolicyRequirementProviderv1ConditionalAccessPolicyIdRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationRequirementPolicyRequirementProvider(input)
	return &out, nil
}
