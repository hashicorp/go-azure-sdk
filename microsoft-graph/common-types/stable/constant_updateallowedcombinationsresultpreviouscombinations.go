package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAllowedCombinationsResultPreviousCombinations string

const (
	UpdateAllowedCombinationsResultPreviousCombinations_DeviceBasedPush             UpdateAllowedCombinationsResultPreviousCombinations = "deviceBasedPush"
	UpdateAllowedCombinationsResultPreviousCombinations_Email                       UpdateAllowedCombinationsResultPreviousCombinations = "email"
	UpdateAllowedCombinationsResultPreviousCombinations_FederatedMultiFactor        UpdateAllowedCombinationsResultPreviousCombinations = "federatedMultiFactor"
	UpdateAllowedCombinationsResultPreviousCombinations_FederatedSingleFactor       UpdateAllowedCombinationsResultPreviousCombinations = "federatedSingleFactor"
	UpdateAllowedCombinationsResultPreviousCombinations_Fido2                       UpdateAllowedCombinationsResultPreviousCombinations = "fido2"
	UpdateAllowedCombinationsResultPreviousCombinations_HardwareOath                UpdateAllowedCombinationsResultPreviousCombinations = "hardwareOath"
	UpdateAllowedCombinationsResultPreviousCombinations_MicrosoftAuthenticatorPush  UpdateAllowedCombinationsResultPreviousCombinations = "microsoftAuthenticatorPush"
	UpdateAllowedCombinationsResultPreviousCombinations_Password                    UpdateAllowedCombinationsResultPreviousCombinations = "password"
	UpdateAllowedCombinationsResultPreviousCombinations_Sms                         UpdateAllowedCombinationsResultPreviousCombinations = "sms"
	UpdateAllowedCombinationsResultPreviousCombinations_SoftwareOath                UpdateAllowedCombinationsResultPreviousCombinations = "softwareOath"
	UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassMultiUse UpdateAllowedCombinationsResultPreviousCombinations = "temporaryAccessPassMultiUse"
	UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassOneTime  UpdateAllowedCombinationsResultPreviousCombinations = "temporaryAccessPassOneTime"
	UpdateAllowedCombinationsResultPreviousCombinations_Voice                       UpdateAllowedCombinationsResultPreviousCombinations = "voice"
	UpdateAllowedCombinationsResultPreviousCombinations_WindowsHelloForBusiness     UpdateAllowedCombinationsResultPreviousCombinations = "windowsHelloForBusiness"
	UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateMultiFactor  UpdateAllowedCombinationsResultPreviousCombinations = "x509CertificateMultiFactor"
	UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateSingleFactor UpdateAllowedCombinationsResultPreviousCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForUpdateAllowedCombinationsResultPreviousCombinations() []string {
	return []string{
		string(UpdateAllowedCombinationsResultPreviousCombinations_DeviceBasedPush),
		string(UpdateAllowedCombinationsResultPreviousCombinations_Email),
		string(UpdateAllowedCombinationsResultPreviousCombinations_FederatedMultiFactor),
		string(UpdateAllowedCombinationsResultPreviousCombinations_FederatedSingleFactor),
		string(UpdateAllowedCombinationsResultPreviousCombinations_Fido2),
		string(UpdateAllowedCombinationsResultPreviousCombinations_HardwareOath),
		string(UpdateAllowedCombinationsResultPreviousCombinations_MicrosoftAuthenticatorPush),
		string(UpdateAllowedCombinationsResultPreviousCombinations_Password),
		string(UpdateAllowedCombinationsResultPreviousCombinations_Sms),
		string(UpdateAllowedCombinationsResultPreviousCombinations_SoftwareOath),
		string(UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassMultiUse),
		string(UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassOneTime),
		string(UpdateAllowedCombinationsResultPreviousCombinations_Voice),
		string(UpdateAllowedCombinationsResultPreviousCombinations_WindowsHelloForBusiness),
		string(UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateMultiFactor),
		string(UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateSingleFactor),
	}
}

func (s *UpdateAllowedCombinationsResultPreviousCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateAllowedCombinationsResultPreviousCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateAllowedCombinationsResultPreviousCombinations(input string) (*UpdateAllowedCombinationsResultPreviousCombinations, error) {
	vals := map[string]UpdateAllowedCombinationsResultPreviousCombinations{
		"devicebasedpush":             UpdateAllowedCombinationsResultPreviousCombinations_DeviceBasedPush,
		"email":                       UpdateAllowedCombinationsResultPreviousCombinations_Email,
		"federatedmultifactor":        UpdateAllowedCombinationsResultPreviousCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       UpdateAllowedCombinationsResultPreviousCombinations_FederatedSingleFactor,
		"fido2":                       UpdateAllowedCombinationsResultPreviousCombinations_Fido2,
		"hardwareoath":                UpdateAllowedCombinationsResultPreviousCombinations_HardwareOath,
		"microsoftauthenticatorpush":  UpdateAllowedCombinationsResultPreviousCombinations_MicrosoftAuthenticatorPush,
		"password":                    UpdateAllowedCombinationsResultPreviousCombinations_Password,
		"sms":                         UpdateAllowedCombinationsResultPreviousCombinations_Sms,
		"softwareoath":                UpdateAllowedCombinationsResultPreviousCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  UpdateAllowedCombinationsResultPreviousCombinations_TemporaryAccessPassOneTime,
		"voice":                       UpdateAllowedCombinationsResultPreviousCombinations_Voice,
		"windowshelloforbusiness":     UpdateAllowedCombinationsResultPreviousCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": UpdateAllowedCombinationsResultPreviousCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateAllowedCombinationsResultPreviousCombinations(input)
	return &out, nil
}
