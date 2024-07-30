package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAllowedCombinationsResultCurrentCombinations string

const (
	UpdateAllowedCombinationsResultCurrentCombinations_DeviceBasedPush             UpdateAllowedCombinationsResultCurrentCombinations = "deviceBasedPush"
	UpdateAllowedCombinationsResultCurrentCombinations_Email                       UpdateAllowedCombinationsResultCurrentCombinations = "email"
	UpdateAllowedCombinationsResultCurrentCombinations_FederatedMultiFactor        UpdateAllowedCombinationsResultCurrentCombinations = "federatedMultiFactor"
	UpdateAllowedCombinationsResultCurrentCombinations_FederatedSingleFactor       UpdateAllowedCombinationsResultCurrentCombinations = "federatedSingleFactor"
	UpdateAllowedCombinationsResultCurrentCombinations_Fido2                       UpdateAllowedCombinationsResultCurrentCombinations = "fido2"
	UpdateAllowedCombinationsResultCurrentCombinations_HardwareOath                UpdateAllowedCombinationsResultCurrentCombinations = "hardwareOath"
	UpdateAllowedCombinationsResultCurrentCombinations_MicrosoftAuthenticatorPush  UpdateAllowedCombinationsResultCurrentCombinations = "microsoftAuthenticatorPush"
	UpdateAllowedCombinationsResultCurrentCombinations_Password                    UpdateAllowedCombinationsResultCurrentCombinations = "password"
	UpdateAllowedCombinationsResultCurrentCombinations_Sms                         UpdateAllowedCombinationsResultCurrentCombinations = "sms"
	UpdateAllowedCombinationsResultCurrentCombinations_SoftwareOath                UpdateAllowedCombinationsResultCurrentCombinations = "softwareOath"
	UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassMultiUse UpdateAllowedCombinationsResultCurrentCombinations = "temporaryAccessPassMultiUse"
	UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassOneTime  UpdateAllowedCombinationsResultCurrentCombinations = "temporaryAccessPassOneTime"
	UpdateAllowedCombinationsResultCurrentCombinations_Voice                       UpdateAllowedCombinationsResultCurrentCombinations = "voice"
	UpdateAllowedCombinationsResultCurrentCombinations_WindowsHelloForBusiness     UpdateAllowedCombinationsResultCurrentCombinations = "windowsHelloForBusiness"
	UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateMultiFactor  UpdateAllowedCombinationsResultCurrentCombinations = "x509CertificateMultiFactor"
	UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateSingleFactor UpdateAllowedCombinationsResultCurrentCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForUpdateAllowedCombinationsResultCurrentCombinations() []string {
	return []string{
		string(UpdateAllowedCombinationsResultCurrentCombinations_DeviceBasedPush),
		string(UpdateAllowedCombinationsResultCurrentCombinations_Email),
		string(UpdateAllowedCombinationsResultCurrentCombinations_FederatedMultiFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinations_FederatedSingleFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinations_Fido2),
		string(UpdateAllowedCombinationsResultCurrentCombinations_HardwareOath),
		string(UpdateAllowedCombinationsResultCurrentCombinations_MicrosoftAuthenticatorPush),
		string(UpdateAllowedCombinationsResultCurrentCombinations_Password),
		string(UpdateAllowedCombinationsResultCurrentCombinations_Sms),
		string(UpdateAllowedCombinationsResultCurrentCombinations_SoftwareOath),
		string(UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassMultiUse),
		string(UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassOneTime),
		string(UpdateAllowedCombinationsResultCurrentCombinations_Voice),
		string(UpdateAllowedCombinationsResultCurrentCombinations_WindowsHelloForBusiness),
		string(UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateMultiFactor),
		string(UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateSingleFactor),
	}
}

func (s *UpdateAllowedCombinationsResultCurrentCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateAllowedCombinationsResultCurrentCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateAllowedCombinationsResultCurrentCombinations(input string) (*UpdateAllowedCombinationsResultCurrentCombinations, error) {
	vals := map[string]UpdateAllowedCombinationsResultCurrentCombinations{
		"devicebasedpush":             UpdateAllowedCombinationsResultCurrentCombinations_DeviceBasedPush,
		"email":                       UpdateAllowedCombinationsResultCurrentCombinations_Email,
		"federatedmultifactor":        UpdateAllowedCombinationsResultCurrentCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       UpdateAllowedCombinationsResultCurrentCombinations_FederatedSingleFactor,
		"fido2":                       UpdateAllowedCombinationsResultCurrentCombinations_Fido2,
		"hardwareoath":                UpdateAllowedCombinationsResultCurrentCombinations_HardwareOath,
		"microsoftauthenticatorpush":  UpdateAllowedCombinationsResultCurrentCombinations_MicrosoftAuthenticatorPush,
		"password":                    UpdateAllowedCombinationsResultCurrentCombinations_Password,
		"sms":                         UpdateAllowedCombinationsResultCurrentCombinations_Sms,
		"softwareoath":                UpdateAllowedCombinationsResultCurrentCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  UpdateAllowedCombinationsResultCurrentCombinations_TemporaryAccessPassOneTime,
		"voice":                       UpdateAllowedCombinationsResultCurrentCombinations_Voice,
		"windowshelloforbusiness":     UpdateAllowedCombinationsResultCurrentCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": UpdateAllowedCombinationsResultCurrentCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateAllowedCombinationsResultCurrentCombinations(input)
	return &out, nil
}
