package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2CombinationConfigurationAppliesToCombinations string

const (
	Fido2CombinationConfigurationAppliesToCombinations_DeviceBasedPush             Fido2CombinationConfigurationAppliesToCombinations = "deviceBasedPush"
	Fido2CombinationConfigurationAppliesToCombinations_Email                       Fido2CombinationConfigurationAppliesToCombinations = "email"
	Fido2CombinationConfigurationAppliesToCombinations_FederatedMultiFactor        Fido2CombinationConfigurationAppliesToCombinations = "federatedMultiFactor"
	Fido2CombinationConfigurationAppliesToCombinations_FederatedSingleFactor       Fido2CombinationConfigurationAppliesToCombinations = "federatedSingleFactor"
	Fido2CombinationConfigurationAppliesToCombinations_Fido2                       Fido2CombinationConfigurationAppliesToCombinations = "fido2"
	Fido2CombinationConfigurationAppliesToCombinations_HardwareOath                Fido2CombinationConfigurationAppliesToCombinations = "hardwareOath"
	Fido2CombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush  Fido2CombinationConfigurationAppliesToCombinations = "microsoftAuthenticatorPush"
	Fido2CombinationConfigurationAppliesToCombinations_Password                    Fido2CombinationConfigurationAppliesToCombinations = "password"
	Fido2CombinationConfigurationAppliesToCombinations_Sms                         Fido2CombinationConfigurationAppliesToCombinations = "sms"
	Fido2CombinationConfigurationAppliesToCombinations_SoftwareOath                Fido2CombinationConfigurationAppliesToCombinations = "softwareOath"
	Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse Fido2CombinationConfigurationAppliesToCombinations = "temporaryAccessPassMultiUse"
	Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime  Fido2CombinationConfigurationAppliesToCombinations = "temporaryAccessPassOneTime"
	Fido2CombinationConfigurationAppliesToCombinations_Voice                       Fido2CombinationConfigurationAppliesToCombinations = "voice"
	Fido2CombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness     Fido2CombinationConfigurationAppliesToCombinations = "windowsHelloForBusiness"
	Fido2CombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor  Fido2CombinationConfigurationAppliesToCombinations = "x509CertificateMultiFactor"
	Fido2CombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor Fido2CombinationConfigurationAppliesToCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForFido2CombinationConfigurationAppliesToCombinations() []string {
	return []string{
		string(Fido2CombinationConfigurationAppliesToCombinations_DeviceBasedPush),
		string(Fido2CombinationConfigurationAppliesToCombinations_Email),
		string(Fido2CombinationConfigurationAppliesToCombinations_FederatedMultiFactor),
		string(Fido2CombinationConfigurationAppliesToCombinations_FederatedSingleFactor),
		string(Fido2CombinationConfigurationAppliesToCombinations_Fido2),
		string(Fido2CombinationConfigurationAppliesToCombinations_HardwareOath),
		string(Fido2CombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush),
		string(Fido2CombinationConfigurationAppliesToCombinations_Password),
		string(Fido2CombinationConfigurationAppliesToCombinations_Sms),
		string(Fido2CombinationConfigurationAppliesToCombinations_SoftwareOath),
		string(Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse),
		string(Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime),
		string(Fido2CombinationConfigurationAppliesToCombinations_Voice),
		string(Fido2CombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness),
		string(Fido2CombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor),
		string(Fido2CombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor),
	}
}

func (s *Fido2CombinationConfigurationAppliesToCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2CombinationConfigurationAppliesToCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2CombinationConfigurationAppliesToCombinations(input string) (*Fido2CombinationConfigurationAppliesToCombinations, error) {
	vals := map[string]Fido2CombinationConfigurationAppliesToCombinations{
		"devicebasedpush":             Fido2CombinationConfigurationAppliesToCombinations_DeviceBasedPush,
		"email":                       Fido2CombinationConfigurationAppliesToCombinations_Email,
		"federatedmultifactor":        Fido2CombinationConfigurationAppliesToCombinations_FederatedMultiFactor,
		"federatedsinglefactor":       Fido2CombinationConfigurationAppliesToCombinations_FederatedSingleFactor,
		"fido2":                       Fido2CombinationConfigurationAppliesToCombinations_Fido2,
		"hardwareoath":                Fido2CombinationConfigurationAppliesToCombinations_HardwareOath,
		"microsoftauthenticatorpush":  Fido2CombinationConfigurationAppliesToCombinations_MicrosoftAuthenticatorPush,
		"password":                    Fido2CombinationConfigurationAppliesToCombinations_Password,
		"sms":                         Fido2CombinationConfigurationAppliesToCombinations_Sms,
		"softwareoath":                Fido2CombinationConfigurationAppliesToCombinations_SoftwareOath,
		"temporaryaccesspassmultiuse": Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  Fido2CombinationConfigurationAppliesToCombinations_TemporaryAccessPassOneTime,
		"voice":                       Fido2CombinationConfigurationAppliesToCombinations_Voice,
		"windowshelloforbusiness":     Fido2CombinationConfigurationAppliesToCombinations_WindowsHelloForBusiness,
		"x509certificatemultifactor":  Fido2CombinationConfigurationAppliesToCombinations_X509CertificateMultiFactor,
		"x509certificatesinglefactor": Fido2CombinationConfigurationAppliesToCombinations_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2CombinationConfigurationAppliesToCombinations(input)
	return &out, nil
}
