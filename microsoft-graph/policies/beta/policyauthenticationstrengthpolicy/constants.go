package policyauthenticationstrengthpolicy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations string

const (
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsdeviceBasedPush             UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "DeviceBasedPush"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsemail                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "Email"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedMultiFactor        UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "FederatedMultiFactor"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedSingleFactor       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "FederatedSingleFactor"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfido2                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "Fido2"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationshardwareOath                UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "HardwareOath"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsmicrosoftAuthenticatorPush  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "MicrosoftAuthenticatorPush"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationspassword                    UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "Password"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssms                         UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "Sms"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssoftwareOath                UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "SoftwareOath"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassMultiUse UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "TemporaryAccessPassMultiUse"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassOneTime  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "TemporaryAccessPassOneTime"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsvoice                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "Voice"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationswindowsHelloForBusiness     UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "WindowsHelloForBusiness"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateMultiFactor  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "X509CertificateMultiFactor"
	UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateSingleFactor UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations = "X509CertificateSingleFactor"
)

func PossibleValuesForUpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations() []string {
	return []string{
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsdeviceBasedPush),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsemail),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedMultiFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedSingleFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfido2),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationshardwareOath),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsmicrosoftAuthenticatorPush),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationspassword),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssms),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssoftwareOath),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassMultiUse),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassOneTime),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsvoice),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationswindowsHelloForBusiness),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateMultiFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateSingleFactor),
	}
}

func (s *UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations(input string) (*UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations, error) {
	vals := map[string]UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations{
		"devicebasedpush":             UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsdeviceBasedPush,
		"email":                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsemail,
		"federatedmultifactor":        UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedMultiFactor,
		"federatedsinglefactor":       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfederatedSingleFactor,
		"fido2":                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsfido2,
		"hardwareoath":                UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationshardwareOath,
		"microsoftauthenticatorpush":  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsmicrosoftAuthenticatorPush,
		"password":                    UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationspassword,
		"sms":                         UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssms,
		"softwareoath":                UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationssoftwareOath,
		"temporaryaccesspassmultiuse": UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationstemporaryAccessPassOneTime,
		"voice":                       UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsvoice,
		"windowshelloforbusiness":     UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationswindowsHelloForBusiness,
		"x509certificatemultifactor":  UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateMultiFactor,
		"x509certificatesinglefactor": UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinationsx509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequestAllowedCombinations(input)
	return &out, nil
}
