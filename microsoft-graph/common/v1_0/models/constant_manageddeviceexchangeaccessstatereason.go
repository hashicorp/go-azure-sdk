package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceExchangeAccessStateReason string

const (
	ManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy ManagedDeviceExchangeAccessStateReason = "AzureADBlockDueToAccessPolicy"
	ManagedDeviceExchangeAccessStateReasoncompliant                     ManagedDeviceExchangeAccessStateReason = "Compliant"
	ManagedDeviceExchangeAccessStateReasoncompromisedPassword           ManagedDeviceExchangeAccessStateReason = "CompromisedPassword"
	ManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp  ManagedDeviceExchangeAccessStateReason = "DeviceNotKnownWithManagedApp"
	ManagedDeviceExchangeAccessStateReasonexchangeDeviceRule            ManagedDeviceExchangeAccessStateReason = "ExchangeDeviceRule"
	ManagedDeviceExchangeAccessStateReasonexchangeGlobalRule            ManagedDeviceExchangeAccessStateReason = "ExchangeGlobalRule"
	ManagedDeviceExchangeAccessStateReasonexchangeIndividualRule        ManagedDeviceExchangeAccessStateReason = "ExchangeIndividualRule"
	ManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy         ManagedDeviceExchangeAccessStateReason = "ExchangeMailboxPolicy"
	ManagedDeviceExchangeAccessStateReasonexchangeUpgrade               ManagedDeviceExchangeAccessStateReason = "ExchangeUpgrade"
	ManagedDeviceExchangeAccessStateReasonmfaRequired                   ManagedDeviceExchangeAccessStateReason = "MfaRequired"
	ManagedDeviceExchangeAccessStateReasonnone                          ManagedDeviceExchangeAccessStateReason = "None"
	ManagedDeviceExchangeAccessStateReasonnotCompliant                  ManagedDeviceExchangeAccessStateReason = "NotCompliant"
	ManagedDeviceExchangeAccessStateReasonnotEnrolled                   ManagedDeviceExchangeAccessStateReason = "NotEnrolled"
	ManagedDeviceExchangeAccessStateReasonother                         ManagedDeviceExchangeAccessStateReason = "Other"
	ManagedDeviceExchangeAccessStateReasonunknown                       ManagedDeviceExchangeAccessStateReason = "Unknown"
	ManagedDeviceExchangeAccessStateReasonunknownLocation               ManagedDeviceExchangeAccessStateReason = "UnknownLocation"
)

func PossibleValuesForManagedDeviceExchangeAccessStateReason() []string {
	return []string{
		string(ManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy),
		string(ManagedDeviceExchangeAccessStateReasoncompliant),
		string(ManagedDeviceExchangeAccessStateReasoncompromisedPassword),
		string(ManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp),
		string(ManagedDeviceExchangeAccessStateReasonexchangeDeviceRule),
		string(ManagedDeviceExchangeAccessStateReasonexchangeGlobalRule),
		string(ManagedDeviceExchangeAccessStateReasonexchangeIndividualRule),
		string(ManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy),
		string(ManagedDeviceExchangeAccessStateReasonexchangeUpgrade),
		string(ManagedDeviceExchangeAccessStateReasonmfaRequired),
		string(ManagedDeviceExchangeAccessStateReasonnone),
		string(ManagedDeviceExchangeAccessStateReasonnotCompliant),
		string(ManagedDeviceExchangeAccessStateReasonnotEnrolled),
		string(ManagedDeviceExchangeAccessStateReasonother),
		string(ManagedDeviceExchangeAccessStateReasonunknown),
		string(ManagedDeviceExchangeAccessStateReasonunknownLocation),
	}
}

func parseManagedDeviceExchangeAccessStateReason(input string) (*ManagedDeviceExchangeAccessStateReason, error) {
	vals := map[string]ManagedDeviceExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": ManagedDeviceExchangeAccessStateReasonazureADBlockDueToAccessPolicy,
		"compliant":                     ManagedDeviceExchangeAccessStateReasoncompliant,
		"compromisedpassword":           ManagedDeviceExchangeAccessStateReasoncompromisedPassword,
		"devicenotknownwithmanagedapp":  ManagedDeviceExchangeAccessStateReasondeviceNotKnownWithManagedApp,
		"exchangedevicerule":            ManagedDeviceExchangeAccessStateReasonexchangeDeviceRule,
		"exchangeglobalrule":            ManagedDeviceExchangeAccessStateReasonexchangeGlobalRule,
		"exchangeindividualrule":        ManagedDeviceExchangeAccessStateReasonexchangeIndividualRule,
		"exchangemailboxpolicy":         ManagedDeviceExchangeAccessStateReasonexchangeMailboxPolicy,
		"exchangeupgrade":               ManagedDeviceExchangeAccessStateReasonexchangeUpgrade,
		"mfarequired":                   ManagedDeviceExchangeAccessStateReasonmfaRequired,
		"none":                          ManagedDeviceExchangeAccessStateReasonnone,
		"notcompliant":                  ManagedDeviceExchangeAccessStateReasonnotCompliant,
		"notenrolled":                   ManagedDeviceExchangeAccessStateReasonnotEnrolled,
		"other":                         ManagedDeviceExchangeAccessStateReasonother,
		"unknown":                       ManagedDeviceExchangeAccessStateReasonunknown,
		"unknownlocation":               ManagedDeviceExchangeAccessStateReasonunknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessStateReason(input)
	return &out, nil
}
