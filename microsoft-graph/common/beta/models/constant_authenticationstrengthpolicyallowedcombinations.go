package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyAllowedCombinations string

const (
	AuthenticationStrengthPolicyAllowedCombinationsdeviceBasedPush             AuthenticationStrengthPolicyAllowedCombinations = "DeviceBasedPush"
	AuthenticationStrengthPolicyAllowedCombinationsemail                       AuthenticationStrengthPolicyAllowedCombinations = "Email"
	AuthenticationStrengthPolicyAllowedCombinationsfederatedMultiFactor        AuthenticationStrengthPolicyAllowedCombinations = "FederatedMultiFactor"
	AuthenticationStrengthPolicyAllowedCombinationsfederatedSingleFactor       AuthenticationStrengthPolicyAllowedCombinations = "FederatedSingleFactor"
	AuthenticationStrengthPolicyAllowedCombinationsfido2                       AuthenticationStrengthPolicyAllowedCombinations = "Fido2"
	AuthenticationStrengthPolicyAllowedCombinationshardwareOath                AuthenticationStrengthPolicyAllowedCombinations = "HardwareOath"
	AuthenticationStrengthPolicyAllowedCombinationsmicrosoftAuthenticatorPush  AuthenticationStrengthPolicyAllowedCombinations = "MicrosoftAuthenticatorPush"
	AuthenticationStrengthPolicyAllowedCombinationspassword                    AuthenticationStrengthPolicyAllowedCombinations = "Password"
	AuthenticationStrengthPolicyAllowedCombinationssms                         AuthenticationStrengthPolicyAllowedCombinations = "Sms"
	AuthenticationStrengthPolicyAllowedCombinationssoftwareOath                AuthenticationStrengthPolicyAllowedCombinations = "SoftwareOath"
	AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassMultiUse AuthenticationStrengthPolicyAllowedCombinations = "TemporaryAccessPassMultiUse"
	AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassOneTime  AuthenticationStrengthPolicyAllowedCombinations = "TemporaryAccessPassOneTime"
	AuthenticationStrengthPolicyAllowedCombinationsvoice                       AuthenticationStrengthPolicyAllowedCombinations = "Voice"
	AuthenticationStrengthPolicyAllowedCombinationswindowsHelloForBusiness     AuthenticationStrengthPolicyAllowedCombinations = "WindowsHelloForBusiness"
	AuthenticationStrengthPolicyAllowedCombinationsx509CertificateMultiFactor  AuthenticationStrengthPolicyAllowedCombinations = "X509CertificateMultiFactor"
	AuthenticationStrengthPolicyAllowedCombinationsx509CertificateSingleFactor AuthenticationStrengthPolicyAllowedCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationStrengthPolicyAllowedCombinations() []string {
	return []string{
		string(AuthenticationStrengthPolicyAllowedCombinationsdeviceBasedPush),
		string(AuthenticationStrengthPolicyAllowedCombinationsemail),
		string(AuthenticationStrengthPolicyAllowedCombinationsfederatedMultiFactor),
		string(AuthenticationStrengthPolicyAllowedCombinationsfederatedSingleFactor),
		string(AuthenticationStrengthPolicyAllowedCombinationsfido2),
		string(AuthenticationStrengthPolicyAllowedCombinationshardwareOath),
		string(AuthenticationStrengthPolicyAllowedCombinationsmicrosoftAuthenticatorPush),
		string(AuthenticationStrengthPolicyAllowedCombinationspassword),
		string(AuthenticationStrengthPolicyAllowedCombinationssms),
		string(AuthenticationStrengthPolicyAllowedCombinationssoftwareOath),
		string(AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassMultiUse),
		string(AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassOneTime),
		string(AuthenticationStrengthPolicyAllowedCombinationsvoice),
		string(AuthenticationStrengthPolicyAllowedCombinationswindowsHelloForBusiness),
		string(AuthenticationStrengthPolicyAllowedCombinationsx509CertificateMultiFactor),
		string(AuthenticationStrengthPolicyAllowedCombinationsx509CertificateSingleFactor),
	}
}

func parseAuthenticationStrengthPolicyAllowedCombinations(input string) (*AuthenticationStrengthPolicyAllowedCombinations, error) {
	vals := map[string]AuthenticationStrengthPolicyAllowedCombinations{
		"devicebasedpush":             AuthenticationStrengthPolicyAllowedCombinationsdeviceBasedPush,
		"email":                       AuthenticationStrengthPolicyAllowedCombinationsemail,
		"federatedmultifactor":        AuthenticationStrengthPolicyAllowedCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationStrengthPolicyAllowedCombinationsfederatedSingleFactor,
		"fido2":                       AuthenticationStrengthPolicyAllowedCombinationsfido2,
		"hardwareoath":                AuthenticationStrengthPolicyAllowedCombinationshardwareOath,
		"microsoftauthenticatorpush":  AuthenticationStrengthPolicyAllowedCombinationsmicrosoftAuthenticatorPush,
		"password":                    AuthenticationStrengthPolicyAllowedCombinationspassword,
		"sms":                         AuthenticationStrengthPolicyAllowedCombinationssms,
		"softwareoath":                AuthenticationStrengthPolicyAllowedCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationStrengthPolicyAllowedCombinationstemporaryAccessPassOneTime,
		"voice":                       AuthenticationStrengthPolicyAllowedCombinationsvoice,
		"windowshelloforbusiness":     AuthenticationStrengthPolicyAllowedCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationStrengthPolicyAllowedCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationStrengthPolicyAllowedCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyAllowedCombinations(input)
	return &out, nil
}
