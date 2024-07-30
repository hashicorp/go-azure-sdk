package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRootCombinations string

const (
	AuthenticationStrengthRootCombinations_DeviceBasedPush             AuthenticationStrengthRootCombinations = "deviceBasedPush"
	AuthenticationStrengthRootCombinations_Email                       AuthenticationStrengthRootCombinations = "email"
	AuthenticationStrengthRootCombinations_FederatedMultiFactor        AuthenticationStrengthRootCombinations = "federatedMultiFactor"
	AuthenticationStrengthRootCombinations_FederatedSingleFactor       AuthenticationStrengthRootCombinations = "federatedSingleFactor"
	AuthenticationStrengthRootCombinations_Fido2                       AuthenticationStrengthRootCombinations = "fido2"
	AuthenticationStrengthRootCombinations_HardwareOath                AuthenticationStrengthRootCombinations = "hardwareOath"
	AuthenticationStrengthRootCombinations_MicrosoftAuthenticatorPush  AuthenticationStrengthRootCombinations = "microsoftAuthenticatorPush"
	AuthenticationStrengthRootCombinations_Password                    AuthenticationStrengthRootCombinations = "password"
	AuthenticationStrengthRootCombinations_Sms                         AuthenticationStrengthRootCombinations = "sms"
	AuthenticationStrengthRootCombinations_SoftwareOath                AuthenticationStrengthRootCombinations = "softwareOath"
	AuthenticationStrengthRootCombinations_TemporaryAccessPassMultiUse AuthenticationStrengthRootCombinations = "temporaryAccessPassMultiUse"
	AuthenticationStrengthRootCombinations_TemporaryAccessPassOneTime  AuthenticationStrengthRootCombinations = "temporaryAccessPassOneTime"
	AuthenticationStrengthRootCombinations_Voice                       AuthenticationStrengthRootCombinations = "voice"
	AuthenticationStrengthRootCombinations_WindowsHelloForBusiness     AuthenticationStrengthRootCombinations = "windowsHelloForBusiness"
	AuthenticationStrengthRootCombinations_X509CertificateMultiFactor  AuthenticationStrengthRootCombinations = "x509CertificateMultiFactor"
	AuthenticationStrengthRootCombinations_X509CertificateSingleFactor AuthenticationStrengthRootCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthRootCombinations() []string {
	return []string{
		string(AuthenticationStrengthRootCombinations_DeviceBasedPush),
		string(AuthenticationStrengthRootCombinations_Email),
		string(AuthenticationStrengthRootCombinations_FederatedMultiFactor),
		string(AuthenticationStrengthRootCombinations_FederatedSingleFactor),
		string(AuthenticationStrengthRootCombinations_Fido2),
		string(AuthenticationStrengthRootCombinations_HardwareOath),
		string(AuthenticationStrengthRootCombinations_MicrosoftAuthenticatorPush),
		string(AuthenticationStrengthRootCombinations_Password),
		string(AuthenticationStrengthRootCombinations_Sms),
		string(AuthenticationStrengthRootCombinations_SoftwareOath),
		string(AuthenticationStrengthRootCombinations_TemporaryAccessPassMultiUse),
		string(AuthenticationStrengthRootCombinations_TemporaryAccessPassOneTime),
		string(AuthenticationStrengthRootCombinations_Voice),
		string(AuthenticationStrengthRootCombinations_WindowsHelloForBusiness),
		string(AuthenticationStrengthRootCombinations_X509CertificateMultiFactor),
		string(AuthenticationStrengthRootCombinations_X509CertificateSingleFactor),
	}
}

func (s *AuthenticationStrengthRootCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthRootCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthRootCombinations(input string) (*AuthenticationStrengthRootCombinations, error) {
	vals := map[string]AuthenticationStrengthRootCombinations{
		"devicebasedpush":             AuthenticationStrengthRootCombinations_DeviceBasedPush,
		"email":                       AuthenticationStrengthRootCombinations_Email,
		"federatedmultifactor":        AuthenticationStrengthRootCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthRootCombinations_FederatedSingleFactor,
		"fido2":                       AuthenticationStrengthRootCombinations_Fido2,
		"hardwareoath":                AuthenticationStrengthRootCombinations_HardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthRootCombinations_MicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthRootCombinations_Password,
		"sms":                         AuthenticationStrengthRootCombinations_Sms,
		"softwareoath":                AuthenticationStrengthRootCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthRootCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthRootCombinations_TemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthRootCombinations_Voice,
		"windowshelloforbusiness":     AuthenticationStrengthRootCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthRootCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthRootCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthRootCombinations(input)
	return &out, nil
}
