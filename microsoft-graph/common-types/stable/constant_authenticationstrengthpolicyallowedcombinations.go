package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyAllowedCombinations string

const (
	AuthenticationStrengthPolicyAllowedCombinations_DeviceBasedPush             AuthenticationStrengthPolicyAllowedCombinations = "deviceBasedPush"
	AuthenticationStrengthPolicyAllowedCombinations_Email                       AuthenticationStrengthPolicyAllowedCombinations = "email"
	AuthenticationStrengthPolicyAllowedCombinations_FederatedMultiFactor        AuthenticationStrengthPolicyAllowedCombinations = "federatedMultiFactor"
	AuthenticationStrengthPolicyAllowedCombinations_FederatedSingleFactor       AuthenticationStrengthPolicyAllowedCombinations = "federatedSingleFactor"
	AuthenticationStrengthPolicyAllowedCombinations_Fido2                       AuthenticationStrengthPolicyAllowedCombinations = "fido2"
	AuthenticationStrengthPolicyAllowedCombinations_HardwareOath                AuthenticationStrengthPolicyAllowedCombinations = "hardwareOath"
	AuthenticationStrengthPolicyAllowedCombinations_MicrosoftAuthenticatorPush  AuthenticationStrengthPolicyAllowedCombinations = "microsoftAuthenticatorPush"
	AuthenticationStrengthPolicyAllowedCombinations_Password                    AuthenticationStrengthPolicyAllowedCombinations = "password"
	AuthenticationStrengthPolicyAllowedCombinations_Sms                         AuthenticationStrengthPolicyAllowedCombinations = "sms"
	AuthenticationStrengthPolicyAllowedCombinations_SoftwareOath                AuthenticationStrengthPolicyAllowedCombinations = "softwareOath"
	AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassMultiUse AuthenticationStrengthPolicyAllowedCombinations = "temporaryAccessPassMultiUse"
	AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassOneTime  AuthenticationStrengthPolicyAllowedCombinations = "temporaryAccessPassOneTime"
	AuthenticationStrengthPolicyAllowedCombinations_Voice                       AuthenticationStrengthPolicyAllowedCombinations = "voice"
	AuthenticationStrengthPolicyAllowedCombinations_WindowsHelloForBusiness     AuthenticationStrengthPolicyAllowedCombinations = "windowsHelloForBusiness"
	AuthenticationStrengthPolicyAllowedCombinations_X509CertificateMultiFactor  AuthenticationStrengthPolicyAllowedCombinations = "x509CertificateMultiFactor"
	AuthenticationStrengthPolicyAllowedCombinations_X509CertificateSingleFactor AuthenticationStrengthPolicyAllowedCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthPolicyAllowedCombinations() []string {
	return []string{
		string(AuthenticationStrengthPolicyAllowedCombinations_DeviceBasedPush),
		string(AuthenticationStrengthPolicyAllowedCombinations_Email),
		string(AuthenticationStrengthPolicyAllowedCombinations_FederatedMultiFactor),
		string(AuthenticationStrengthPolicyAllowedCombinations_FederatedSingleFactor),
		string(AuthenticationStrengthPolicyAllowedCombinations_Fido2),
		string(AuthenticationStrengthPolicyAllowedCombinations_HardwareOath),
		string(AuthenticationStrengthPolicyAllowedCombinations_MicrosoftAuthenticatorPush),
		string(AuthenticationStrengthPolicyAllowedCombinations_Password),
		string(AuthenticationStrengthPolicyAllowedCombinations_Sms),
		string(AuthenticationStrengthPolicyAllowedCombinations_SoftwareOath),
		string(AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassMultiUse),
		string(AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassOneTime),
		string(AuthenticationStrengthPolicyAllowedCombinations_Voice),
		string(AuthenticationStrengthPolicyAllowedCombinations_WindowsHelloForBusiness),
		string(AuthenticationStrengthPolicyAllowedCombinations_X509CertificateMultiFactor),
		string(AuthenticationStrengthPolicyAllowedCombinations_X509CertificateSingleFactor),
	}
}

func (s *AuthenticationStrengthPolicyAllowedCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthPolicyAllowedCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthPolicyAllowedCombinations(input string) (*AuthenticationStrengthPolicyAllowedCombinations, error) {
	vals := map[string]AuthenticationStrengthPolicyAllowedCombinations{
		"devicebasedpush":             AuthenticationStrengthPolicyAllowedCombinations_DeviceBasedPush,
		"email":                       AuthenticationStrengthPolicyAllowedCombinations_Email,
		"federatedmultifactor":        AuthenticationStrengthPolicyAllowedCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthPolicyAllowedCombinations_FederatedSingleFactor,
		"fido2":                       AuthenticationStrengthPolicyAllowedCombinations_Fido2,
		"hardwareoath":                AuthenticationStrengthPolicyAllowedCombinations_HardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthPolicyAllowedCombinations_MicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthPolicyAllowedCombinations_Password,
		"sms":                         AuthenticationStrengthPolicyAllowedCombinations_Sms,
		"softwareoath":                AuthenticationStrengthPolicyAllowedCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthPolicyAllowedCombinations_TemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthPolicyAllowedCombinations_Voice,
		"windowshelloforbusiness":     AuthenticationStrengthPolicyAllowedCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthPolicyAllowedCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthPolicyAllowedCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyAllowedCombinations(input)
	return &out, nil
}
