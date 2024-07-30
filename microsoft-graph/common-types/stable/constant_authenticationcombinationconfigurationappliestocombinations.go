package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationCombinationConfigurationAppliesToCombinations string

const (
	AuthenticationCombinationConfigurationAppliesToCombinations_DeviceBasedPush             AuthenticationCombinationConfigurationAppliesToCombinations = "deviceBasedPush"
	AuthenticationCombinationConfigurationAppliesToCombinations_Email                       AuthenticationCombinationConfigurationAppliesToCombinations = "email"
	AuthenticationCombinationConfigurationAppliesToCombinations_FederatedMultiFactor        AuthenticationCombinationConfigurationAppliesToCombinations = "federatedMultiFactor"
	AuthenticationCombinationConfigurationAppliesToCombinations_FederatedSingleFactor       AuthenticationCombinationConfigurationAppliesToCombinations = "federatedSingleFactor"
	AuthenticationCombinationConfigurationAppliesToCombinations_Fido2                       AuthenticationCombinationConfigurationAppliesToCombinations = "fido2"
	AuthenticationCombinationConfigurationAppliesToCombinations_HardwareOath                AuthenticationCombinationConfigurationAppliesToCombinations = "hardwareOath"
	AuthenticationCombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush  AuthenticationCombinationConfigurationAppliesToCombinations = "microsoftAuthenticatorPush"
	AuthenticationCombinationConfigurationAppliesToCombinations_Password                    AuthenticationCombinationConfigurationAppliesToCombinations = "password"
	AuthenticationCombinationConfigurationAppliesToCombinations_Sms                         AuthenticationCombinationConfigurationAppliesToCombinations = "sms"
	AuthenticationCombinationConfigurationAppliesToCombinations_SoftwareOath                AuthenticationCombinationConfigurationAppliesToCombinations = "softwareOath"
	AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse AuthenticationCombinationConfigurationAppliesToCombinations = "temporaryAccessPassMultiUse"
	AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime  AuthenticationCombinationConfigurationAppliesToCombinations = "temporaryAccessPassOneTime"
	AuthenticationCombinationConfigurationAppliesToCombinations_Voice                       AuthenticationCombinationConfigurationAppliesToCombinations = "voice"
	AuthenticationCombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness     AuthenticationCombinationConfigurationAppliesToCombinations = "windowsHelloForBusiness"
	AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor  AuthenticationCombinationConfigurationAppliesToCombinations = "x509CertificateMultiFactor"
	AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor AuthenticationCombinationConfigurationAppliesToCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationCombinationConfigurationAppliesToCombinations() []string {
	return []string{
		string(AuthenticationCombinationConfigurationAppliesToCombinations_DeviceBasedPush),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_Email),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_FederatedMultiFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_FederatedSingleFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_Fido2),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_HardwareOath),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_Password),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_Sms),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_SoftwareOath),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_Voice),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor),
	}
}

func (s *AuthenticationCombinationConfigurationAppliesToCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationCombinationConfigurationAppliesToCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationCombinationConfigurationAppliesToCombinations(input string) (*AuthenticationCombinationConfigurationAppliesToCombinations, error) {
	vals := map[string]AuthenticationCombinationConfigurationAppliesToCombinations{
		"devicebasedpush":             AuthenticationCombinationConfigurationAppliesToCombinations_DeviceBasedPush,
		"email":                       AuthenticationCombinationConfigurationAppliesToCombinations_Email,
		"federatedmultifactor":        AuthenticationCombinationConfigurationAppliesToCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationCombinationConfigurationAppliesToCombinations_FederatedSingleFactor,
		"fido2":                       AuthenticationCombinationConfigurationAppliesToCombinations_Fido2,
		"hardwareoath":                AuthenticationCombinationConfigurationAppliesToCombinations_HardwareOath,
		"microsoftauthenticatorpush":  AuthenticationCombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush,
		"password":                    AuthenticationCombinationConfigurationAppliesToCombinations_Password,
		"sms":                         AuthenticationCombinationConfigurationAppliesToCombinations_Sms,
		"softwareoath":                AuthenticationCombinationConfigurationAppliesToCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationCombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime,
		"voice":                       AuthenticationCombinationConfigurationAppliesToCombinations_Voice,
		"windowshelloforbusiness":     AuthenticationCombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationCombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationCombinationConfigurationAppliesToCombinations(input)
	return &out, nil
}
