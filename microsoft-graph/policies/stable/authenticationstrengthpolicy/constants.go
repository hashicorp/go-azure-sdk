package authenticationstrengthpolicy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations string

const (
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsDeviceBasedPush             UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "deviceBasedPush"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsEmail                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "email"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedMultiFactor        UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "federatedMultiFactor"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedSingleFactor       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "federatedSingleFactor"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFido2                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "fido2"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsHardwareOath                UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "hardwareOath"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsMicrosoftAuthenticatorPush  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "microsoftAuthenticatorPush"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsPassword                    UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "password"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSms                         UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "sms"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSoftwareOath                UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "softwareOath"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassMultiUse UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "temporaryAccessPassMultiUse"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassOneTime  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "temporaryAccessPassOneTime"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsVoice                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "voice"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsWindowsHelloForBusiness     UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "windowsHelloForBusiness"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateMultiFactor  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "x509CertificateMultiFactor"
	UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateSingleFactor UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations = "x509CertificateSingleFactor"
)

func PossibleValuesForUpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations() []string {
	return []string{
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsDeviceBasedPush),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsEmail),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedMultiFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedSingleFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFido2),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsHardwareOath),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsMicrosoftAuthenticatorPush),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsPassword),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSms),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSoftwareOath),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassMultiUse),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassOneTime),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsVoice),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsWindowsHelloForBusiness),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateMultiFactor),
		string(UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateSingleFactor),
	}
}

func (s *UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations(input string) (*UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations, error) {
	vals := map[string]UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations{
		"devicebasedpush":             UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsDeviceBasedPush,
		"email":                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsEmail,
		"federatedmultifactor":        UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedMultiFactor,
		"federatedsinglefactor":       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFederatedSingleFactor,
		"fido2":                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsFido2,
		"hardwareoath":                UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsHardwareOath,
		"microsoftauthenticatorpush":  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsMicrosoftAuthenticatorPush,
		"password":                    UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsPassword,
		"sms":                         UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSms,
		"softwareoath":                UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsSoftwareOath,
		"temporaryaccesspassmultiuse": UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsTemporaryAccessPassOneTime,
		"voice":                       UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsVoice,
		"windowshelloforbusiness":     UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsWindowsHelloForBusiness,
		"x509certificatemultifactor":  UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateMultiFactor,
		"x509certificatesinglefactor": UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinationsX509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequestAllowedCombinations(input)
	return &out, nil
}
