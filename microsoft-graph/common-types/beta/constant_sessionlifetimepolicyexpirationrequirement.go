package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionLifetimePolicyExpirationRequirement string

const (
	SessionLifetimePolicyExpirationRequirement_AudienceTokenLifetimePolicy                       SessionLifetimePolicyExpirationRequirement = "audienceTokenLifetimePolicy"
	SessionLifetimePolicyExpirationRequirement_NgcMfa                                            SessionLifetimePolicyExpirationRequirement = "ngcMfa"
	SessionLifetimePolicyExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices SessionLifetimePolicyExpirationRequirement = "rememberMultifactorAuthenticationOnTrustedDevices"
	SessionLifetimePolicyExpirationRequirement_SignInFrequencyEveryTime                          SessionLifetimePolicyExpirationRequirement = "signInFrequencyEveryTime"
	SessionLifetimePolicyExpirationRequirement_SignInFrequencyPeriodicReauthentication           SessionLifetimePolicyExpirationRequirement = "signInFrequencyPeriodicReauthentication"
	SessionLifetimePolicyExpirationRequirement_TenantTokenLifetimePolicy                         SessionLifetimePolicyExpirationRequirement = "tenantTokenLifetimePolicy"
)

func PossibleValuesForSessionLifetimePolicyExpirationRequirement() []string {
	return []string{
		string(SessionLifetimePolicyExpirationRequirement_AudienceTokenLifetimePolicy),
		string(SessionLifetimePolicyExpirationRequirement_NgcMfa),
		string(SessionLifetimePolicyExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices),
		string(SessionLifetimePolicyExpirationRequirement_SignInFrequencyEveryTime),
		string(SessionLifetimePolicyExpirationRequirement_SignInFrequencyPeriodicReauthentication),
		string(SessionLifetimePolicyExpirationRequirement_TenantTokenLifetimePolicy),
	}
}

func (s *SessionLifetimePolicyExpirationRequirement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSessionLifetimePolicyExpirationRequirement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSessionLifetimePolicyExpirationRequirement(input string) (*SessionLifetimePolicyExpirationRequirement, error) {
	vals := map[string]SessionLifetimePolicyExpirationRequirement{
		"audiencetokenlifetimepolicy": SessionLifetimePolicyExpirationRequirement_AudienceTokenLifetimePolicy,
		"ngcmfa":                      SessionLifetimePolicyExpirationRequirement_NgcMfa,
		"remembermultifactorauthenticationontrusteddevices": SessionLifetimePolicyExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices,
		"signinfrequencyeverytime":                          SessionLifetimePolicyExpirationRequirement_SignInFrequencyEveryTime,
		"signinfrequencyperiodicreauthentication":           SessionLifetimePolicyExpirationRequirement_SignInFrequencyPeriodicReauthentication,
		"tenanttokenlifetimepolicy":                         SessionLifetimePolicyExpirationRequirement_TenantTokenLifetimePolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionLifetimePolicyExpirationRequirement(input)
	return &out, nil
}
