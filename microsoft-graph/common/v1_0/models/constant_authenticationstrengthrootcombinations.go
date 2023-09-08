package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRootCombinations string

const (
	AuthenticationStrengthRootCombinationsdeviceBasedPush             AuthenticationStrengthRootCombinations = "DeviceBasedPush"
	AuthenticationStrengthRootCombinationsemail                       AuthenticationStrengthRootCombinations = "Email"
	AuthenticationStrengthRootCombinationsfederatedMultiFactor        AuthenticationStrengthRootCombinations = "FederatedMultiFactor"
	AuthenticationStrengthRootCombinationsfederatedSingleFactor       AuthenticationStrengthRootCombinations = "FederatedSingleFactor"
	AuthenticationStrengthRootCombinationsfido2                       AuthenticationStrengthRootCombinations = "Fido2"
	AuthenticationStrengthRootCombinationshardwareOath                AuthenticationStrengthRootCombinations = "HardwareOath"
	AuthenticationStrengthRootCombinationsmicrosoftAuthenticatorPush  AuthenticationStrengthRootCombinations = "MicrosoftAuthenticatorPush"
	AuthenticationStrengthRootCombinationspassword                    AuthenticationStrengthRootCombinations = "Password"
	AuthenticationStrengthRootCombinationssms                         AuthenticationStrengthRootCombinations = "Sms"
	AuthenticationStrengthRootCombinationssoftwareOath                AuthenticationStrengthRootCombinations = "SoftwareOath"
	AuthenticationStrengthRootCombinationstemporaryAccessPassMultiUse AuthenticationStrengthRootCombinations = "TemporaryAccessPassMultiUse"
	AuthenticationStrengthRootCombinationstemporaryAccessPassOneTime  AuthenticationStrengthRootCombinations = "TemporaryAccessPassOneTime"
	AuthenticationStrengthRootCombinationsvoice                       AuthenticationStrengthRootCombinations = "Voice"
	AuthenticationStrengthRootCombinationswindowsHelloForBusiness     AuthenticationStrengthRootCombinations = "WindowsHelloForBusiness"
	AuthenticationStrengthRootCombinationsx509CertificateMultiFactor  AuthenticationStrengthRootCombinations = "X509CertificateMultiFactor"
	AuthenticationStrengthRootCombinationsx509CertificateSingleFactor AuthenticationStrengthRootCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthRootCombinations() []string {
	return []string{
		string(AuthenticationStrengthRootCombinationsdeviceBasedPush),
		string(AuthenticationStrengthRootCombinationsemail),
		string(AuthenticationStrengthRootCombinationsfederatedMultiFactor),
		string(AuthenticationStrengthRootCombinationsfederatedSingleFactor),
		string(AuthenticationStrengthRootCombinationsfido2),
		string(AuthenticationStrengthRootCombinationshardwareOath),
		string(AuthenticationStrengthRootCombinationsmicrosoftAuthenticatorPush),
		string(AuthenticationStrengthRootCombinationspassword),
		string(AuthenticationStrengthRootCombinationssms),
		string(AuthenticationStrengthRootCombinationssoftwareOath),
		string(AuthenticationStrengthRootCombinationstemporaryAccessPassMultiUse),
		string(AuthenticationStrengthRootCombinationstemporaryAccessPassOneTime),
		string(AuthenticationStrengthRootCombinationsvoice),
		string(AuthenticationStrengthRootCombinationswindowsHelloForBusiness),
		string(AuthenticationStrengthRootCombinationsx509CertificateMultiFactor),
		string(AuthenticationStrengthRootCombinationsx509CertificateSingleFactor),
	}
}

func parseAuthenticationStrengthRootCombinations(input string) (*AuthenticationStrengthRootCombinations, error) {
	vals := map[string]AuthenticationStrengthRootCombinations{
		"devicebasedpush":             AuthenticationStrengthRootCombinationsdeviceBasedPush,
		"email":                       AuthenticationStrengthRootCombinationsemail,
		"federatedmultifactor":        AuthenticationStrengthRootCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthRootCombinationsfederatedSingleFactor,
		"fido2":                       AuthenticationStrengthRootCombinationsfido2,
		"hardwareoath":                AuthenticationStrengthRootCombinationshardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthRootCombinationsmicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthRootCombinationspassword,
		"sms":                         AuthenticationStrengthRootCombinationssms,
		"softwareoath":                AuthenticationStrengthRootCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthRootCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthRootCombinationstemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthRootCombinationsvoice,
		"windowshelloforbusiness":     AuthenticationStrengthRootCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthRootCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthRootCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthRootCombinations(input)
	return &out, nil
}
