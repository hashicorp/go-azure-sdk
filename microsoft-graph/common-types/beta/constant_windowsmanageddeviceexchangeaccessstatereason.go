package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceExchangeAccessStateReason string

const (
	WindowsManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy WindowsManagedDeviceExchangeAccessStateReason = "azureADBlockDueToAccessPolicy"
	WindowsManagedDeviceExchangeAccessStateReason_Compliant                     WindowsManagedDeviceExchangeAccessStateReason = "compliant"
	WindowsManagedDeviceExchangeAccessStateReason_CompromisedPassword           WindowsManagedDeviceExchangeAccessStateReason = "compromisedPassword"
	WindowsManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp  WindowsManagedDeviceExchangeAccessStateReason = "deviceNotKnownWithManagedApp"
	WindowsManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule            WindowsManagedDeviceExchangeAccessStateReason = "exchangeDeviceRule"
	WindowsManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule            WindowsManagedDeviceExchangeAccessStateReason = "exchangeGlobalRule"
	WindowsManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule        WindowsManagedDeviceExchangeAccessStateReason = "exchangeIndividualRule"
	WindowsManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy         WindowsManagedDeviceExchangeAccessStateReason = "exchangeMailboxPolicy"
	WindowsManagedDeviceExchangeAccessStateReason_ExchangeUpgrade               WindowsManagedDeviceExchangeAccessStateReason = "exchangeUpgrade"
	WindowsManagedDeviceExchangeAccessStateReason_MfaRequired                   WindowsManagedDeviceExchangeAccessStateReason = "mfaRequired"
	WindowsManagedDeviceExchangeAccessStateReason_None                          WindowsManagedDeviceExchangeAccessStateReason = "none"
	WindowsManagedDeviceExchangeAccessStateReason_NotCompliant                  WindowsManagedDeviceExchangeAccessStateReason = "notCompliant"
	WindowsManagedDeviceExchangeAccessStateReason_NotEnrolled                   WindowsManagedDeviceExchangeAccessStateReason = "notEnrolled"
	WindowsManagedDeviceExchangeAccessStateReason_Other                         WindowsManagedDeviceExchangeAccessStateReason = "other"
	WindowsManagedDeviceExchangeAccessStateReason_Unknown                       WindowsManagedDeviceExchangeAccessStateReason = "unknown"
	WindowsManagedDeviceExchangeAccessStateReason_UnknownLocation               WindowsManagedDeviceExchangeAccessStateReason = "unknownLocation"
)

func PossibleValuesForWindowsManagedDeviceExchangeAccessStateReason() []string {
	return []string{
		string(WindowsManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy),
		string(WindowsManagedDeviceExchangeAccessStateReason_Compliant),
		string(WindowsManagedDeviceExchangeAccessStateReason_CompromisedPassword),
		string(WindowsManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp),
		string(WindowsManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule),
		string(WindowsManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule),
		string(WindowsManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule),
		string(WindowsManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy),
		string(WindowsManagedDeviceExchangeAccessStateReason_ExchangeUpgrade),
		string(WindowsManagedDeviceExchangeAccessStateReason_MfaRequired),
		string(WindowsManagedDeviceExchangeAccessStateReason_None),
		string(WindowsManagedDeviceExchangeAccessStateReason_NotCompliant),
		string(WindowsManagedDeviceExchangeAccessStateReason_NotEnrolled),
		string(WindowsManagedDeviceExchangeAccessStateReason_Other),
		string(WindowsManagedDeviceExchangeAccessStateReason_Unknown),
		string(WindowsManagedDeviceExchangeAccessStateReason_UnknownLocation),
	}
}

func (s *WindowsManagedDeviceExchangeAccessStateReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceExchangeAccessStateReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceExchangeAccessStateReason(input string) (*WindowsManagedDeviceExchangeAccessStateReason, error) {
	vals := map[string]WindowsManagedDeviceExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": WindowsManagedDeviceExchangeAccessStateReason_AzureADBlockDueToAccessPolicy,
		"compliant":                     WindowsManagedDeviceExchangeAccessStateReason_Compliant,
		"compromisedpassword":           WindowsManagedDeviceExchangeAccessStateReason_CompromisedPassword,
		"devicenotknownwithmanagedapp":  WindowsManagedDeviceExchangeAccessStateReason_DeviceNotKnownWithManagedApp,
		"exchangedevicerule":            WindowsManagedDeviceExchangeAccessStateReason_ExchangeDeviceRule,
		"exchangeglobalrule":            WindowsManagedDeviceExchangeAccessStateReason_ExchangeGlobalRule,
		"exchangeindividualrule":        WindowsManagedDeviceExchangeAccessStateReason_ExchangeIndividualRule,
		"exchangemailboxpolicy":         WindowsManagedDeviceExchangeAccessStateReason_ExchangeMailboxPolicy,
		"exchangeupgrade":               WindowsManagedDeviceExchangeAccessStateReason_ExchangeUpgrade,
		"mfarequired":                   WindowsManagedDeviceExchangeAccessStateReason_MfaRequired,
		"none":                          WindowsManagedDeviceExchangeAccessStateReason_None,
		"notcompliant":                  WindowsManagedDeviceExchangeAccessStateReason_NotCompliant,
		"notenrolled":                   WindowsManagedDeviceExchangeAccessStateReason_NotEnrolled,
		"other":                         WindowsManagedDeviceExchangeAccessStateReason_Other,
		"unknown":                       WindowsManagedDeviceExchangeAccessStateReason_Unknown,
		"unknownlocation":               WindowsManagedDeviceExchangeAccessStateReason_UnknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceExchangeAccessStateReason(input)
	return &out, nil
}
