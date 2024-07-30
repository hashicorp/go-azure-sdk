package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRootAuthenticationCombinations string

const (
	AuthenticationStrengthRootAuthenticationCombinations_DeviceBasedPush             AuthenticationStrengthRootAuthenticationCombinations = "deviceBasedPush"
	AuthenticationStrengthRootAuthenticationCombinations_Email                       AuthenticationStrengthRootAuthenticationCombinations = "email"
	AuthenticationStrengthRootAuthenticationCombinations_FederatedMultiFactor        AuthenticationStrengthRootAuthenticationCombinations = "federatedMultiFactor"
	AuthenticationStrengthRootAuthenticationCombinations_FederatedSingleFactor       AuthenticationStrengthRootAuthenticationCombinations = "federatedSingleFactor"
	AuthenticationStrengthRootAuthenticationCombinations_Fido2                       AuthenticationStrengthRootAuthenticationCombinations = "fido2"
	AuthenticationStrengthRootAuthenticationCombinations_HardwareOath                AuthenticationStrengthRootAuthenticationCombinations = "hardwareOath"
	AuthenticationStrengthRootAuthenticationCombinations_MicrosoftAuthenticatorPush  AuthenticationStrengthRootAuthenticationCombinations = "microsoftAuthenticatorPush"
	AuthenticationStrengthRootAuthenticationCombinations_Password                    AuthenticationStrengthRootAuthenticationCombinations = "password"
	AuthenticationStrengthRootAuthenticationCombinations_Sms                         AuthenticationStrengthRootAuthenticationCombinations = "sms"
	AuthenticationStrengthRootAuthenticationCombinations_SoftwareOath                AuthenticationStrengthRootAuthenticationCombinations = "softwareOath"
	AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassMultiUse AuthenticationStrengthRootAuthenticationCombinations = "temporaryAccessPassMultiUse"
	AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassOneTime  AuthenticationStrengthRootAuthenticationCombinations = "temporaryAccessPassOneTime"
	AuthenticationStrengthRootAuthenticationCombinations_Voice                       AuthenticationStrengthRootAuthenticationCombinations = "voice"
	AuthenticationStrengthRootAuthenticationCombinations_WindowsHelloForBusiness     AuthenticationStrengthRootAuthenticationCombinations = "windowsHelloForBusiness"
	AuthenticationStrengthRootAuthenticationCombinations_X509CertificateMultiFactor  AuthenticationStrengthRootAuthenticationCombinations = "x509CertificateMultiFactor"
	AuthenticationStrengthRootAuthenticationCombinations_X509CertificateSingleFactor AuthenticationStrengthRootAuthenticationCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthRootAuthenticationCombinations() []string {
	return []string{
		string(AuthenticationStrengthRootAuthenticationCombinations_DeviceBasedPush),
		string(AuthenticationStrengthRootAuthenticationCombinations_Email),
		string(AuthenticationStrengthRootAuthenticationCombinations_FederatedMultiFactor),
		string(AuthenticationStrengthRootAuthenticationCombinations_FederatedSingleFactor),
		string(AuthenticationStrengthRootAuthenticationCombinations_Fido2),
		string(AuthenticationStrengthRootAuthenticationCombinations_HardwareOath),
		string(AuthenticationStrengthRootAuthenticationCombinations_MicrosoftAuthenticatorPush),
		string(AuthenticationStrengthRootAuthenticationCombinations_Password),
		string(AuthenticationStrengthRootAuthenticationCombinations_Sms),
		string(AuthenticationStrengthRootAuthenticationCombinations_SoftwareOath),
		string(AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassMultiUse),
		string(AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassOneTime),
		string(AuthenticationStrengthRootAuthenticationCombinations_Voice),
		string(AuthenticationStrengthRootAuthenticationCombinations_WindowsHelloForBusiness),
		string(AuthenticationStrengthRootAuthenticationCombinations_X509CertificateMultiFactor),
		string(AuthenticationStrengthRootAuthenticationCombinations_X509CertificateSingleFactor),
	}
}

func (s *AuthenticationStrengthRootAuthenticationCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthRootAuthenticationCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthRootAuthenticationCombinations(input string) (*AuthenticationStrengthRootAuthenticationCombinations, error) {
	vals := map[string]AuthenticationStrengthRootAuthenticationCombinations{
		"devicebasedpush":             AuthenticationStrengthRootAuthenticationCombinations_DeviceBasedPush,
		"email":                       AuthenticationStrengthRootAuthenticationCombinations_Email,
		"federatedmultifactor":        AuthenticationStrengthRootAuthenticationCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthRootAuthenticationCombinations_FederatedSingleFactor,
		"fido2":                       AuthenticationStrengthRootAuthenticationCombinations_Fido2,
		"hardwareoath":                AuthenticationStrengthRootAuthenticationCombinations_HardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthRootAuthenticationCombinations_MicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthRootAuthenticationCombinations_Password,
		"sms":                         AuthenticationStrengthRootAuthenticationCombinations_Sms,
		"softwareoath":                AuthenticationStrengthRootAuthenticationCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthRootAuthenticationCombinations_TemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthRootAuthenticationCombinations_Voice,
		"windowshelloforbusiness":     AuthenticationStrengthRootAuthenticationCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthRootAuthenticationCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthRootAuthenticationCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthRootAuthenticationCombinations(input)
	return &out, nil
}
