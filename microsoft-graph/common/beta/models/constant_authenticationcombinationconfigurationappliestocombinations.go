package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationCombinationConfigurationAppliesToCombinations string

const (
	AuthenticationCombinationConfigurationAppliesToCombinationsdeviceBasedPush             AuthenticationCombinationConfigurationAppliesToCombinations = "DeviceBasedPush"
	AuthenticationCombinationConfigurationAppliesToCombinationsemail                       AuthenticationCombinationConfigurationAppliesToCombinations = "Email"
	AuthenticationCombinationConfigurationAppliesToCombinationsfederatedMultiFactor        AuthenticationCombinationConfigurationAppliesToCombinations = "FederatedMultiFactor"
	AuthenticationCombinationConfigurationAppliesToCombinationsfederatedSingleFactor       AuthenticationCombinationConfigurationAppliesToCombinations = "FederatedSingleFactor"
	AuthenticationCombinationConfigurationAppliesToCombinationsfido2                       AuthenticationCombinationConfigurationAppliesToCombinations = "Fido2"
	AuthenticationCombinationConfigurationAppliesToCombinationshardwareOath                AuthenticationCombinationConfigurationAppliesToCombinations = "HardwareOath"
	AuthenticationCombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush  AuthenticationCombinationConfigurationAppliesToCombinations = "MicrosoftAuthenticatorPush"
	AuthenticationCombinationConfigurationAppliesToCombinationspassword                    AuthenticationCombinationConfigurationAppliesToCombinations = "Password"
	AuthenticationCombinationConfigurationAppliesToCombinationssms                         AuthenticationCombinationConfigurationAppliesToCombinations = "Sms"
	AuthenticationCombinationConfigurationAppliesToCombinationssoftwareOath                AuthenticationCombinationConfigurationAppliesToCombinations = "SoftwareOath"
	AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse AuthenticationCombinationConfigurationAppliesToCombinations = "TemporaryAccessPassMultiUse"
	AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime  AuthenticationCombinationConfigurationAppliesToCombinations = "TemporaryAccessPassOneTime"
	AuthenticationCombinationConfigurationAppliesToCombinationsvoice                       AuthenticationCombinationConfigurationAppliesToCombinations = "Voice"
	AuthenticationCombinationConfigurationAppliesToCombinationswindowsHelloForBusiness     AuthenticationCombinationConfigurationAppliesToCombinations = "WindowsHelloForBusiness"
	AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor  AuthenticationCombinationConfigurationAppliesToCombinations = "X509CertificateMultiFactor"
	AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor AuthenticationCombinationConfigurationAppliesToCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationCombinationConfigurationAppliesToCombinations() []string {
	return []string{
		string(AuthenticationCombinationConfigurationAppliesToCombinationsdeviceBasedPush),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsemail),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsfederatedMultiFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsfederatedSingleFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsfido2),
		string(AuthenticationCombinationConfigurationAppliesToCombinationshardwareOath),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush),
		string(AuthenticationCombinationConfigurationAppliesToCombinationspassword),
		string(AuthenticationCombinationConfigurationAppliesToCombinationssms),
		string(AuthenticationCombinationConfigurationAppliesToCombinationssoftwareOath),
		string(AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse),
		string(AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsvoice),
		string(AuthenticationCombinationConfigurationAppliesToCombinationswindowsHelloForBusiness),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor),
		string(AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor),
	}
}

func parseAuthenticationCombinationConfigurationAppliesToCombinations(input string) (*AuthenticationCombinationConfigurationAppliesToCombinations, error) {
	vals := map[string]AuthenticationCombinationConfigurationAppliesToCombinations{
		"devicebasedpush":             AuthenticationCombinationConfigurationAppliesToCombinationsdeviceBasedPush,
		"email":                       AuthenticationCombinationConfigurationAppliesToCombinationsemail,
		"federatedmultifactor":        AuthenticationCombinationConfigurationAppliesToCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationCombinationConfigurationAppliesToCombinationsfederatedSingleFactor,
		"fido2":                       AuthenticationCombinationConfigurationAppliesToCombinationsfido2,
		"hardwareoath":                AuthenticationCombinationConfigurationAppliesToCombinationshardwareOath,
		"microsoftauthenticatorpush":  AuthenticationCombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush,
		"password":                    AuthenticationCombinationConfigurationAppliesToCombinationspassword,
		"sms":                         AuthenticationCombinationConfigurationAppliesToCombinationssms,
		"softwareoath":                AuthenticationCombinationConfigurationAppliesToCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationCombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime,
		"voice":                       AuthenticationCombinationConfigurationAppliesToCombinationsvoice,
		"windowshelloforbusiness":     AuthenticationCombinationConfigurationAppliesToCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationCombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationCombinationConfigurationAppliesToCombinations(input)
	return &out, nil
}
