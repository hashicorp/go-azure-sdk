package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2CombinationConfigurationAppliesToCombinations string

const (
	Fido2CombinationConfigurationAppliesToCombinationsdeviceBasedPush             Fido2CombinationConfigurationAppliesToCombinations = "DeviceBasedPush"
	Fido2CombinationConfigurationAppliesToCombinationsemail                       Fido2CombinationConfigurationAppliesToCombinations = "Email"
	Fido2CombinationConfigurationAppliesToCombinationsfederatedMultiFactor        Fido2CombinationConfigurationAppliesToCombinations = "FederatedMultiFactor"
	Fido2CombinationConfigurationAppliesToCombinationsfederatedSingleFactor       Fido2CombinationConfigurationAppliesToCombinations = "FederatedSingleFactor"
	Fido2CombinationConfigurationAppliesToCombinationsfido2                       Fido2CombinationConfigurationAppliesToCombinations = "Fido2"
	Fido2CombinationConfigurationAppliesToCombinationshardwareOath                Fido2CombinationConfigurationAppliesToCombinations = "HardwareOath"
	Fido2CombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush  Fido2CombinationConfigurationAppliesToCombinations = "MicrosoftAuthenticatorPush"
	Fido2CombinationConfigurationAppliesToCombinationspassword                    Fido2CombinationConfigurationAppliesToCombinations = "Password"
	Fido2CombinationConfigurationAppliesToCombinationssms                         Fido2CombinationConfigurationAppliesToCombinations = "Sms"
	Fido2CombinationConfigurationAppliesToCombinationssoftwareOath                Fido2CombinationConfigurationAppliesToCombinations = "SoftwareOath"
	Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse Fido2CombinationConfigurationAppliesToCombinations = "TemporaryAccessPassMultiUse"
	Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime  Fido2CombinationConfigurationAppliesToCombinations = "TemporaryAccessPassOneTime"
	Fido2CombinationConfigurationAppliesToCombinationsvoice                       Fido2CombinationConfigurationAppliesToCombinations = "Voice"
	Fido2CombinationConfigurationAppliesToCombinationswindowsHelloForBusiness     Fido2CombinationConfigurationAppliesToCombinations = "WindowsHelloForBusiness"
	Fido2CombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor  Fido2CombinationConfigurationAppliesToCombinations = "X509CertificateMultiFactor"
	Fido2CombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor Fido2CombinationConfigurationAppliesToCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForFido2CombinationConfigurationAppliesToCombinations() []string {
	return []string{
		string(Fido2CombinationConfigurationAppliesToCombinationsdeviceBasedPush),
		string(Fido2CombinationConfigurationAppliesToCombinationsemail),
		string(Fido2CombinationConfigurationAppliesToCombinationsfederatedMultiFactor),
		string(Fido2CombinationConfigurationAppliesToCombinationsfederatedSingleFactor),
		string(Fido2CombinationConfigurationAppliesToCombinationsfido2),
		string(Fido2CombinationConfigurationAppliesToCombinationshardwareOath),
		string(Fido2CombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush),
		string(Fido2CombinationConfigurationAppliesToCombinationspassword),
		string(Fido2CombinationConfigurationAppliesToCombinationssms),
		string(Fido2CombinationConfigurationAppliesToCombinationssoftwareOath),
		string(Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse),
		string(Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime),
		string(Fido2CombinationConfigurationAppliesToCombinationsvoice),
		string(Fido2CombinationConfigurationAppliesToCombinationswindowsHelloForBusiness),
		string(Fido2CombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor),
		string(Fido2CombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor),
	}
}

func parseFido2CombinationConfigurationAppliesToCombinations(input string) (*Fido2CombinationConfigurationAppliesToCombinations, error) {
	vals := map[string]Fido2CombinationConfigurationAppliesToCombinations{
		"devicebasedpush":             Fido2CombinationConfigurationAppliesToCombinationsdeviceBasedPush,
		"email":                       Fido2CombinationConfigurationAppliesToCombinationsemail,
		"federatedmultifactor":        Fido2CombinationConfigurationAppliesToCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       Fido2CombinationConfigurationAppliesToCombinationsfederatedSingleFactor,
		"fido2":                       Fido2CombinationConfigurationAppliesToCombinationsfido2,
		"hardwareoath":                Fido2CombinationConfigurationAppliesToCombinationshardwareOath,
		"microsoftauthenticatorpush":  Fido2CombinationConfigurationAppliesToCombinationsmicrosoftAuthenticatorPush,
		"password":                    Fido2CombinationConfigurationAppliesToCombinationspassword,
		"sms":                         Fido2CombinationConfigurationAppliesToCombinationssms,
		"softwareoath":                Fido2CombinationConfigurationAppliesToCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  Fido2CombinationConfigurationAppliesToCombinationstemporaryAccessPassOneTime,
		"voice":                       Fido2CombinationConfigurationAppliesToCombinationsvoice,
		"windowshelloforbusiness":     Fido2CombinationConfigurationAppliesToCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  Fido2CombinationConfigurationAppliesToCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": Fido2CombinationConfigurationAppliesToCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2CombinationConfigurationAppliesToCombinations(input)
	return &out, nil
}
