package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionLifetimePolicyExpirationRequirement string

const (
	SessionLifetimePolicyExpirationRequirementaudienceTokenLifetimePolicy                       SessionLifetimePolicyExpirationRequirement = "AudienceTokenLifetimePolicy"
	SessionLifetimePolicyExpirationRequirementngcMfa                                            SessionLifetimePolicyExpirationRequirement = "NgcMfa"
	SessionLifetimePolicyExpirationRequirementrememberMultifactorAuthenticationOnTrustedDevices SessionLifetimePolicyExpirationRequirement = "RememberMultifactorAuthenticationOnTrustedDevices"
	SessionLifetimePolicyExpirationRequirementsignInFrequencyEveryTime                          SessionLifetimePolicyExpirationRequirement = "SignInFrequencyEveryTime"
	SessionLifetimePolicyExpirationRequirementsignInFrequencyPeriodicReauthentication           SessionLifetimePolicyExpirationRequirement = "SignInFrequencyPeriodicReauthentication"
	SessionLifetimePolicyExpirationRequirementtenantTokenLifetimePolicy                         SessionLifetimePolicyExpirationRequirement = "TenantTokenLifetimePolicy"
)

func PossibleValuesForSessionLifetimePolicyExpirationRequirement() []string {
	return []string{
		string(SessionLifetimePolicyExpirationRequirementaudienceTokenLifetimePolicy),
		string(SessionLifetimePolicyExpirationRequirementngcMfa),
		string(SessionLifetimePolicyExpirationRequirementrememberMultifactorAuthenticationOnTrustedDevices),
		string(SessionLifetimePolicyExpirationRequirementsignInFrequencyEveryTime),
		string(SessionLifetimePolicyExpirationRequirementsignInFrequencyPeriodicReauthentication),
		string(SessionLifetimePolicyExpirationRequirementtenantTokenLifetimePolicy),
	}
}

func parseSessionLifetimePolicyExpirationRequirement(input string) (*SessionLifetimePolicyExpirationRequirement, error) {
	vals := map[string]SessionLifetimePolicyExpirationRequirement{
		"audiencetokenlifetimepolicy": SessionLifetimePolicyExpirationRequirementaudienceTokenLifetimePolicy,
		"ngcmfa":                      SessionLifetimePolicyExpirationRequirementngcMfa,
		"remembermultifactorauthenticationontrusteddevices": SessionLifetimePolicyExpirationRequirementrememberMultifactorAuthenticationOnTrustedDevices,
		"signinfrequencyeverytime":                          SessionLifetimePolicyExpirationRequirementsignInFrequencyEveryTime,
		"signinfrequencyperiodicreauthentication":           SessionLifetimePolicyExpirationRequirementsignInFrequencyPeriodicReauthentication,
		"tenanttokenlifetimepolicy":                         SessionLifetimePolicyExpirationRequirementtenantTokenLifetimePolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionLifetimePolicyExpirationRequirement(input)
	return &out, nil
}
