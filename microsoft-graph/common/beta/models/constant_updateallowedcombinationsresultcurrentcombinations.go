package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAllowedCombinationsResultCurrentCombinations string

const (
	UpdateAllowedCombinationsResultCurrentCombinationsdeviceBasedPush             UpdateAllowedCombinationsResultCurrentCombinations = "DeviceBasedPush"
	UpdateAllowedCombinationsResultCurrentCombinationsemail                       UpdateAllowedCombinationsResultCurrentCombinations = "Email"
	UpdateAllowedCombinationsResultCurrentCombinationsfederatedMultiFactor        UpdateAllowedCombinationsResultCurrentCombinations = "FederatedMultiFactor"
	UpdateAllowedCombinationsResultCurrentCombinationsfederatedSingleFactor       UpdateAllowedCombinationsResultCurrentCombinations = "FederatedSingleFactor"
	UpdateAllowedCombinationsResultCurrentCombinationsfido2                       UpdateAllowedCombinationsResultCurrentCombinations = "Fido2"
	UpdateAllowedCombinationsResultCurrentCombinationshardwareOath                UpdateAllowedCombinationsResultCurrentCombinations = "HardwareOath"
	UpdateAllowedCombinationsResultCurrentCombinationsmicrosoftAuthenticatorPush  UpdateAllowedCombinationsResultCurrentCombinations = "MicrosoftAuthenticatorPush"
	UpdateAllowedCombinationsResultCurrentCombinationspassword                    UpdateAllowedCombinationsResultCurrentCombinations = "Password"
	UpdateAllowedCombinationsResultCurrentCombinationssms                         UpdateAllowedCombinationsResultCurrentCombinations = "Sms"
	UpdateAllowedCombinationsResultCurrentCombinationssoftwareOath                UpdateAllowedCombinationsResultCurrentCombinations = "SoftwareOath"
	UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassMultiUse UpdateAllowedCombinationsResultCurrentCombinations = "TemporaryAccessPassMultiUse"
	UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassOneTime  UpdateAllowedCombinationsResultCurrentCombinations = "TemporaryAccessPassOneTime"
	UpdateAllowedCombinationsResultCurrentCombinationsvoice                       UpdateAllowedCombinationsResultCurrentCombinations = "Voice"
	UpdateAllowedCombinationsResultCurrentCombinationswindowsHelloForBusiness     UpdateAllowedCombinationsResultCurrentCombinations = "WindowsHelloForBusiness"
	UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateMultiFactor  UpdateAllowedCombinationsResultCurrentCombinations = "X509CertificateMultiFactor"
	UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateSingleFactor UpdateAllowedCombinationsResultCurrentCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForUpdateAllowedCombinationsResultCurrentCombinations() []string {
	return []string{
		string(UpdateAllowedCombinationsResultCurrentCombinationsdeviceBasedPush),
		string(UpdateAllowedCombinationsResultCurrentCombinationsemail),
		string(UpdateAllowedCombinationsResultCurrentCombinationsfederatedMultiFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinationsfederatedSingleFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinationsfido2),
		string(UpdateAllowedCombinationsResultCurrentCombinationshardwareOath),
		string(UpdateAllowedCombinationsResultCurrentCombinationsmicrosoftAuthenticatorPush),
		string(UpdateAllowedCombinationsResultCurrentCombinationspassword),
		string(UpdateAllowedCombinationsResultCurrentCombinationssms),
		string(UpdateAllowedCombinationsResultCurrentCombinationssoftwareOath),
		string(UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassMultiUse),
		string(UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassOneTime),
		string(UpdateAllowedCombinationsResultCurrentCombinationsvoice),
		string(UpdateAllowedCombinationsResultCurrentCombinationswindowsHelloForBusiness),
		string(UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateMultiFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateSingleFactor),
	}
}

func parseUpdateAllowedCombinationsResultCurrentCombinations(input string) (*UpdateAllowedCombinationsResultCurrentCombinations, error) {
	vals := map[string]UpdateAllowedCombinationsResultCurrentCombinations{
		"devicebasedpush":             UpdateAllowedCombinationsResultCurrentCombinationsdeviceBasedPush,
		"email":                       UpdateAllowedCombinationsResultCurrentCombinationsemail,
		"federatedmultifactor":        UpdateAllowedCombinationsResultCurrentCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       UpdateAllowedCombinationsResultCurrentCombinationsfederatedSingleFactor,
		"fido2":                       UpdateAllowedCombinationsResultCurrentCombinationsfido2,
		"hardwareoath":                UpdateAllowedCombinationsResultCurrentCombinationshardwareOath,
		"microsoftauthenticatorpush":  UpdateAllowedCombinationsResultCurrentCombinationsmicrosoftAuthenticatorPush,
		"password":                    UpdateAllowedCombinationsResultCurrentCombinationspassword,
		"sms":                         UpdateAllowedCombinationsResultCurrentCombinationssms,
		"softwareoath":                UpdateAllowedCombinationsResultCurrentCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  UpdateAllowedCombinationsResultCurrentCombinationstemporaryAccessPassOneTime,
		"voice":                       UpdateAllowedCombinationsResultCurrentCombinationsvoice,
		"windowshelloforbusiness":     UpdateAllowedCombinationsResultCurrentCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": UpdateAllowedCombinationsResultCurrentCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateAllowedCombinationsResultCurrentCombinations(input)
	return &out, nil
}
