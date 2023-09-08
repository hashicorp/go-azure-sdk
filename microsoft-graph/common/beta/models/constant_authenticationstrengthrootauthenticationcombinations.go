package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRootAuthenticationCombinations string

const (
	AuthenticationStrengthRootAuthenticationCombinationsdeviceBasedPush             AuthenticationStrengthRootAuthenticationCombinations = "DeviceBasedPush"
	AuthenticationStrengthRootAuthenticationCombinationsemail                       AuthenticationStrengthRootAuthenticationCombinations = "Email"
	AuthenticationStrengthRootAuthenticationCombinationsfederatedMultiFactor        AuthenticationStrengthRootAuthenticationCombinations = "FederatedMultiFactor"
	AuthenticationStrengthRootAuthenticationCombinationsfederatedSingleFactor       AuthenticationStrengthRootAuthenticationCombinations = "FederatedSingleFactor"
	AuthenticationStrengthRootAuthenticationCombinationsfido2                       AuthenticationStrengthRootAuthenticationCombinations = "Fido2"
	AuthenticationStrengthRootAuthenticationCombinationshardwareOath                AuthenticationStrengthRootAuthenticationCombinations = "HardwareOath"
	AuthenticationStrengthRootAuthenticationCombinationsmicrosoftAuthenticatorPush  AuthenticationStrengthRootAuthenticationCombinations = "MicrosoftAuthenticatorPush"
	AuthenticationStrengthRootAuthenticationCombinationspassword                    AuthenticationStrengthRootAuthenticationCombinations = "Password"
	AuthenticationStrengthRootAuthenticationCombinationssms                         AuthenticationStrengthRootAuthenticationCombinations = "Sms"
	AuthenticationStrengthRootAuthenticationCombinationssoftwareOath                AuthenticationStrengthRootAuthenticationCombinations = "SoftwareOath"
	AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassMultiUse AuthenticationStrengthRootAuthenticationCombinations = "TemporaryAccessPassMultiUse"
	AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassOneTime  AuthenticationStrengthRootAuthenticationCombinations = "TemporaryAccessPassOneTime"
	AuthenticationStrengthRootAuthenticationCombinationsvoice                       AuthenticationStrengthRootAuthenticationCombinations = "Voice"
	AuthenticationStrengthRootAuthenticationCombinationswindowsHelloForBusiness     AuthenticationStrengthRootAuthenticationCombinations = "WindowsHelloForBusiness"
	AuthenticationStrengthRootAuthenticationCombinationsx509CertificateMultiFactor  AuthenticationStrengthRootAuthenticationCombinations = "X509CertificateMultiFactor"
	AuthenticationStrengthRootAuthenticationCombinationsx509CertificateSingleFactor AuthenticationStrengthRootAuthenticationCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthRootAuthenticationCombinations() []string {
	return []string{
		string(AuthenticationStrengthRootAuthenticationCombinationsdeviceBasedPush),
		string(AuthenticationStrengthRootAuthenticationCombinationsemail),
		string(AuthenticationStrengthRootAuthenticationCombinationsfederatedMultiFactor),
		string(AuthenticationStrengthRootAuthenticationCombinationsfederatedSingleFactor),
		string(AuthenticationStrengthRootAuthenticationCombinationsfido2),
		string(AuthenticationStrengthRootAuthenticationCombinationshardwareOath),
		string(AuthenticationStrengthRootAuthenticationCombinationsmicrosoftAuthenticatorPush),
		string(AuthenticationStrengthRootAuthenticationCombinationspassword),
		string(AuthenticationStrengthRootAuthenticationCombinationssms),
		string(AuthenticationStrengthRootAuthenticationCombinationssoftwareOath),
		string(AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassMultiUse),
		string(AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassOneTime),
		string(AuthenticationStrengthRootAuthenticationCombinationsvoice),
		string(AuthenticationStrengthRootAuthenticationCombinationswindowsHelloForBusiness),
		string(AuthenticationStrengthRootAuthenticationCombinationsx509CertificateMultiFactor),
		string(AuthenticationStrengthRootAuthenticationCombinationsx509CertificateSingleFactor),
	}
}

func parseAuthenticationStrengthRootAuthenticationCombinations(input string) (*AuthenticationStrengthRootAuthenticationCombinations, error) {
	vals := map[string]AuthenticationStrengthRootAuthenticationCombinations{
		"devicebasedpush":             AuthenticationStrengthRootAuthenticationCombinationsdeviceBasedPush,
		"email":                       AuthenticationStrengthRootAuthenticationCombinationsemail,
		"federatedmultifactor":        AuthenticationStrengthRootAuthenticationCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthRootAuthenticationCombinationsfederatedSingleFactor,
		"fido2":                       AuthenticationStrengthRootAuthenticationCombinationsfido2,
		"hardwareoath":                AuthenticationStrengthRootAuthenticationCombinationshardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthRootAuthenticationCombinationsmicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthRootAuthenticationCombinationspassword,
		"sms":                         AuthenticationStrengthRootAuthenticationCombinationssms,
		"softwareoath":                AuthenticationStrengthRootAuthenticationCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthRootAuthenticationCombinationstemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthRootAuthenticationCombinationsvoice,
		"windowshelloforbusiness":     AuthenticationStrengthRootAuthenticationCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthRootAuthenticationCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthRootAuthenticationCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthRootAuthenticationCombinations(input)
	return &out, nil
}
