package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationRequirementPolicyRequirementProvider string

const (
	AuthenticationRequirementPolicyRequirementProvider_AccountCompromisePolicies                         AuthenticationRequirementPolicyRequirementProvider = "accountCompromisePolicies"
	AuthenticationRequirementPolicyRequirementProvider_BaselineProtection                                AuthenticationRequirementPolicyRequirementProvider = "baselineProtection"
	AuthenticationRequirementPolicyRequirementProvider_CrossTenantOutboundRule                           AuthenticationRequirementPolicyRequirementProvider = "crossTenantOutboundRule"
	AuthenticationRequirementPolicyRequirementProvider_EnforcedForCspAdmins                              AuthenticationRequirementPolicyRequirementProvider = "enforcedForCspAdmins"
	AuthenticationRequirementPolicyRequirementProvider_GpsLocationCondition                              AuthenticationRequirementPolicyRequirementProvider = "gpsLocationCondition"
	AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByBaselineProtection       AuthenticationRequirementPolicyRequirementProvider = "mfaRegistrationRequiredByBaselineProtection"
	AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy AuthenticationRequirementPolicyRequirementProvider = "mfaRegistrationRequiredByIdentityProtectionPolicy"
	AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess   AuthenticationRequirementPolicyRequirementProvider = "mfaRegistrationRequiredByMultiConditionalAccess"
	AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredBySecurityDefaults         AuthenticationRequirementPolicyRequirementProvider = "mfaRegistrationRequiredBySecurityDefaults"
	AuthenticationRequirementPolicyRequirementProvider_MultiConditionalAccess                            AuthenticationRequirementPolicyRequirementProvider = "multiConditionalAccess"
	AuthenticationRequirementPolicyRequirementProvider_ProofUpCodeRequest                                AuthenticationRequirementPolicyRequirementProvider = "proofUpCodeRequest"
	AuthenticationRequirementPolicyRequirementProvider_Request                                           AuthenticationRequirementPolicyRequirementProvider = "request"
	AuthenticationRequirementPolicyRequirementProvider_RiskBasedPolicy                                   AuthenticationRequirementPolicyRequirementProvider = "riskBasedPolicy"
	AuthenticationRequirementPolicyRequirementProvider_SecurityDefaults                                  AuthenticationRequirementPolicyRequirementProvider = "securityDefaults"
	AuthenticationRequirementPolicyRequirementProvider_ServicePrincipal                                  AuthenticationRequirementPolicyRequirementProvider = "servicePrincipal"
	AuthenticationRequirementPolicyRequirementProvider_TenantSessionRiskPolicy                           AuthenticationRequirementPolicyRequirementProvider = "tenantSessionRiskPolicy"
	AuthenticationRequirementPolicyRequirementProvider_User                                              AuthenticationRequirementPolicyRequirementProvider = "user"
	AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccess                               AuthenticationRequirementPolicyRequirementProvider = "v1ConditionalAccess"
	AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessDependency                     AuthenticationRequirementPolicyRequirementProvider = "v1ConditionalAccessDependency"
	AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessPolicyIdRequested              AuthenticationRequirementPolicyRequirementProvider = "v1ConditionalAccessPolicyIdRequested"
)

func PossibleValuesForAuthenticationRequirementPolicyRequirementProvider() []string {
	return []string{
		string(AuthenticationRequirementPolicyRequirementProvider_AccountCompromisePolicies),
		string(AuthenticationRequirementPolicyRequirementProvider_BaselineProtection),
		string(AuthenticationRequirementPolicyRequirementProvider_CrossTenantOutboundRule),
		string(AuthenticationRequirementPolicyRequirementProvider_EnforcedForCspAdmins),
		string(AuthenticationRequirementPolicyRequirementProvider_GpsLocationCondition),
		string(AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByBaselineProtection),
		string(AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy),
		string(AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredBySecurityDefaults),
		string(AuthenticationRequirementPolicyRequirementProvider_MultiConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProvider_ProofUpCodeRequest),
		string(AuthenticationRequirementPolicyRequirementProvider_Request),
		string(AuthenticationRequirementPolicyRequirementProvider_RiskBasedPolicy),
		string(AuthenticationRequirementPolicyRequirementProvider_SecurityDefaults),
		string(AuthenticationRequirementPolicyRequirementProvider_ServicePrincipal),
		string(AuthenticationRequirementPolicyRequirementProvider_TenantSessionRiskPolicy),
		string(AuthenticationRequirementPolicyRequirementProvider_User),
		string(AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccess),
		string(AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessDependency),
		string(AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessPolicyIdRequested),
	}
}

func (s *AuthenticationRequirementPolicyRequirementProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationRequirementPolicyRequirementProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationRequirementPolicyRequirementProvider(input string) (*AuthenticationRequirementPolicyRequirementProvider, error) {
	vals := map[string]AuthenticationRequirementPolicyRequirementProvider{
		"accountcompromisepolicies":                         AuthenticationRequirementPolicyRequirementProvider_AccountCompromisePolicies,
		"baselineprotection":                                AuthenticationRequirementPolicyRequirementProvider_BaselineProtection,
		"crosstenantoutboundrule":                           AuthenticationRequirementPolicyRequirementProvider_CrossTenantOutboundRule,
		"enforcedforcspadmins":                              AuthenticationRequirementPolicyRequirementProvider_EnforcedForCspAdmins,
		"gpslocationcondition":                              AuthenticationRequirementPolicyRequirementProvider_GpsLocationCondition,
		"mfaregistrationrequiredbybaselineprotection":       AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByBaselineProtection,
		"mfaregistrationrequiredbyidentityprotectionpolicy": AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy,
		"mfaregistrationrequiredbymulticonditionalaccess":   AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess,
		"mfaregistrationrequiredbysecuritydefaults":         AuthenticationRequirementPolicyRequirementProvider_MfaRegistrationRequiredBySecurityDefaults,
		"multiconditionalaccess":                            AuthenticationRequirementPolicyRequirementProvider_MultiConditionalAccess,
		"proofupcoderequest":                                AuthenticationRequirementPolicyRequirementProvider_ProofUpCodeRequest,
		"request":                                           AuthenticationRequirementPolicyRequirementProvider_Request,
		"riskbasedpolicy":                                   AuthenticationRequirementPolicyRequirementProvider_RiskBasedPolicy,
		"securitydefaults":                                  AuthenticationRequirementPolicyRequirementProvider_SecurityDefaults,
		"serviceprincipal":                                  AuthenticationRequirementPolicyRequirementProvider_ServicePrincipal,
		"tenantsessionriskpolicy":                           AuthenticationRequirementPolicyRequirementProvider_TenantSessionRiskPolicy,
		"user":                                              AuthenticationRequirementPolicyRequirementProvider_User,
		"v1conditionalaccess":                               AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccess,
		"v1conditionalaccessdependency":                     AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessDependency,
		"v1conditionalaccesspolicyidrequested":              AuthenticationRequirementPolicyRequirementProvider_V1ConditionalAccessPolicyIdRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationRequirementPolicyRequirementProvider(input)
	return &out, nil
}
