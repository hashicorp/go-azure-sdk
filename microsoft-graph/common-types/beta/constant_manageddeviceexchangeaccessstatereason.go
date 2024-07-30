package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceExchangeAccessStateReason string

const (
	ManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy ManagedDeviceExchangeAccessStateReason = "azureADBlockDueToAccessPolicy"
	ManagedDeviceExchangeAccessStateReason_Compliant                     ManagedDeviceExchangeAccessStateReason = "compliant"
	ManagedDeviceExchangeAccessStateReason_CompromisedPassword           ManagedDeviceExchangeAccessStateReason = "compromisedPassword"
	ManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp  ManagedDeviceExchangeAccessStateReason = "deviceNotKnownWithManagedApp"
	ManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule            ManagedDeviceExchangeAccessStateReason = "exchangeDeviceRule"
	ManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule            ManagedDeviceExchangeAccessStateReason = "exchangeGlobalRule"
	ManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule        ManagedDeviceExchangeAccessStateReason = "exchangeIndividualRule"
	ManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy         ManagedDeviceExchangeAccessStateReason = "exchangeMailboxPolicy"
	ManagedDeviceExchangeAccessStateReason_ExchangeUpgrade               ManagedDeviceExchangeAccessStateReason = "exchangeUpgrade"
	ManagedDeviceExchangeAccessStateReason_MfaRequired                   ManagedDeviceExchangeAccessStateReason = "mfaRequired"
	ManagedDeviceExchangeAccessStateReason_None                          ManagedDeviceExchangeAccessStateReason = "none"
	ManagedDeviceExchangeAccessStateReason_NotCompliant                  ManagedDeviceExchangeAccessStateReason = "notCompliant"
	ManagedDeviceExchangeAccessStateReason_NotEnrolled                   ManagedDeviceExchangeAccessStateReason = "notEnrolled"
	ManagedDeviceExchangeAccessStateReason_Other                         ManagedDeviceExchangeAccessStateReason = "other"
	ManagedDeviceExchangeAccessStateReason_Unknown                       ManagedDeviceExchangeAccessStateReason = "unknown"
	ManagedDeviceExchangeAccessStateReason_UnknownLocation               ManagedDeviceExchangeAccessStateReason = "unknownLocation"
)

func PossibleValuesForManagedDeviceExchangeAccessStateReason() []string {
	return []string{
		string(ManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy),
		string(ManagedDeviceExchangeAccessStateReason_Compliant),
		string(ManagedDeviceExchangeAccessStateReason_CompromisedPassword),
		string(ManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp),
		string(ManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule),
		string(ManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule),
		string(ManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule),
		string(ManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy),
		string(ManagedDeviceExchangeAccessStateReason_ExchangeUpgrade),
		string(ManagedDeviceExchangeAccessStateReason_MfaRequired),
		string(ManagedDeviceExchangeAccessStateReason_None),
		string(ManagedDeviceExchangeAccessStateReason_NotCompliant),
		string(ManagedDeviceExchangeAccessStateReason_NotEnrolled),
		string(ManagedDeviceExchangeAccessStateReason_Other),
		string(ManagedDeviceExchangeAccessStateReason_Unknown),
		string(ManagedDeviceExchangeAccessStateReason_UnknownLocation),
	}
}

func (s *ManagedDeviceExchangeAccessStateReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceExchangeAccessStateReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceExchangeAccessStateReason(input string) (*ManagedDeviceExchangeAccessStateReason, error) {
	vals := map[string]ManagedDeviceExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": ManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy,
		"compliant":                     ManagedDeviceExchangeAccessStateReason_Compliant,
		"compromisedpassword":           ManagedDeviceExchangeAccessStateReason_CompromisedPassword,
		"devicenotknownwithmanagedapp":  ManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp,
		"exchangedevicerule":            ManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule,
		"exchangeglobalrule":            ManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule,
		"exchangeindividualrule":        ManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule,
		"exchangemailboxpolicy":         ManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy,
		"exchangeupgrade":               ManagedDeviceExchangeAccessStateReason_ExchangeUpgrade,
		"mfarequired":                   ManagedDeviceExchangeAccessStateReason_MfaRequired,
		"none":                          ManagedDeviceExchangeAccessStateReason_None,
		"notcompliant":                  ManagedDeviceExchangeAccessStateReason_NotCompliant,
		"notenrolled":                   ManagedDeviceExchangeAccessStateReason_NotEnrolled,
		"other":                         ManagedDeviceExchangeAccessStateReason_Other,
		"unknown":                       ManagedDeviceExchangeAccessStateReason_Unknown,
		"unknownlocation":               ManagedDeviceExchangeAccessStateReason_UnknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessStateReason(input)
	return &out, nil
}
