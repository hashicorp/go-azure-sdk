package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceExchangeAccessStateReason string

const (
	WindowsManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy WindowsManagedDeviceExchangeAccessStateReason = "AzureADBlockDueToAccessPolicy"
	WindowsManagedDeviceExchangeAccessStateReasoncompliant                     WindowsManagedDeviceExchangeAccessStateReason = "Compliant"
	WindowsManagedDeviceExchangeAccessStateReasoncompromisedPassword           WindowsManagedDeviceExchangeAccessStateReason = "CompromisedPassword"
	WindowsManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp  WindowsManagedDeviceExchangeAccessStateReason = "DeviceNotKnownWithManagedApp"
	WindowsManagedDeviceExchangeAccessStateReasonexchangeDeviceRule            WindowsManagedDeviceExchangeAccessStateReason = "ExchangeDeviceRule"
	WindowsManagedDeviceExchangeAccessStateReasonexchangeGlobalRule            WindowsManagedDeviceExchangeAccessStateReason = "ExchangeGlobalRule"
	WindowsManagedDeviceExchangeAccessStateReasonexchangeIndividualRule        WindowsManagedDeviceExchangeAccessStateReason = "ExchangeIndividualRule"
	WindowsManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy         WindowsManagedDeviceExchangeAccessStateReason = "ExchangeMailboxPolicy"
	WindowsManagedDeviceExchangeAccessStateReasonexchangeUpgrade               WindowsManagedDeviceExchangeAccessStateReason = "ExchangeUpgrade"
	WindowsManagedDeviceExchangeAccessStateReasonmfaRequired                   WindowsManagedDeviceExchangeAccessStateReason = "MfaRequired"
	WindowsManagedDeviceExchangeAccessStateReasonnone                          WindowsManagedDeviceExchangeAccessStateReason = "None"
	WindowsManagedDeviceExchangeAccessStateReasonnotCompliant                  WindowsManagedDeviceExchangeAccessStateReason = "NotCompliant"
	WindowsManagedDeviceExchangeAccessStateReasonnotEnrolled                   WindowsManagedDeviceExchangeAccessStateReason = "NotEnrolled"
	WindowsManagedDeviceExchangeAccessStateReasonother                         WindowsManagedDeviceExchangeAccessStateReason = "Other"
	WindowsManagedDeviceExchangeAccessStateReasonunknown                       WindowsManagedDeviceExchangeAccessStateReason = "Unknown"
	WindowsManagedDeviceExchangeAccessStateReasonunknownLocation               WindowsManagedDeviceExchangeAccessStateReason = "UnknownLocation"
)

func PossibleValuesForWindowsManagedDeviceExchangeAccessStateReason() []string {
	return []string{
		string(WindowsManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy),
		string(WindowsManagedDeviceExchangeAccessStateReasoncompliant),
		string(WindowsManagedDeviceExchangeAccessStateReasoncompromisedPassword),
		string(WindowsManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp),
		string(WindowsManagedDeviceExchangeAccessStateReasonexchangeDeviceRule),
		string(WindowsManagedDeviceExchangeAccessStateReasonexchangeGlobalRule),
		string(WindowsManagedDeviceExchangeAccessStateReasonexchangeIndividualRule),
		string(WindowsManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy),
		string(WindowsManagedDeviceExchangeAccessStateReasonexchangeUpgrade),
		string(WindowsManagedDeviceExchangeAccessStateReasonmfaRequired),
		string(WindowsManagedDeviceExchangeAccessStateReasonnone),
		string(WindowsManagedDeviceExchangeAccessStateReasonnotCompliant),
		string(WindowsManagedDeviceExchangeAccessStateReasonnotEnrolled),
		string(WindowsManagedDeviceExchangeAccessStateReasonother),
		string(WindowsManagedDeviceExchangeAccessStateReasonunknown),
		string(WindowsManagedDeviceExchangeAccessStateReasonunknownLocation),
	}
}

func parseWindowsManagedDeviceExchangeAccessStateReason(input string) (*WindowsManagedDeviceExchangeAccessStateReason, error) {
	vals := map[string]WindowsManagedDeviceExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": WindowsManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy,
		"compliant":                     WindowsManagedDeviceExchangeAccessStateReasoncompliant,
		"compromisedpassword":           WindowsManagedDeviceExchangeAccessStateReasoncompromisedPassword,
		"devicenotknownwithmanagedapp":  WindowsManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp,
		"exchangedevicerule":            WindowsManagedDeviceExchangeAccessStateReasonexchangeDeviceRule,
		"exchangeglobalrule":            WindowsManagedDeviceExchangeAccessStateReasonexchangeGlobalRule,
		"exchangeindividualrule":        WindowsManagedDeviceExchangeAccessStateReasonexchangeIndividualRule,
		"exchangemailboxpolicy":         WindowsManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy,
		"exchangeupgrade":               WindowsManagedDeviceExchangeAccessStateReasonexchangeUpgrade,
		"mfarequired":                   WindowsManagedDeviceExchangeAccessStateReasonmfaRequired,
		"none":                          WindowsManagedDeviceExchangeAccessStateReasonnone,
		"notcompliant":                  WindowsManagedDeviceExchangeAccessStateReasonnotCompliant,
		"notenrolled":                   WindowsManagedDeviceExchangeAccessStateReasonnotEnrolled,
		"other":                         WindowsManagedDeviceExchangeAccessStateReasonother,
		"unknown":                       WindowsManagedDeviceExchangeAccessStateReasonunknown,
		"unknownlocation":               WindowsManagedDeviceExchangeAccessStateReasonunknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceExchangeAccessStateReason(input)
	return &out, nil
}
