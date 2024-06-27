package reservations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedScopeType string

const (
	AppliedScopeTypeManagementGroup AppliedScopeType = "ManagementGroup"
	AppliedScopeTypeShared          AppliedScopeType = "Shared"
	AppliedScopeTypeSingle          AppliedScopeType = "Single"
)

func PossibleValuesForAppliedScopeType() []string {
	return []string{
		string(AppliedScopeTypeManagementGroup),
		string(AppliedScopeTypeShared),
		string(AppliedScopeTypeSingle),
	}
}

func (s *AppliedScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppliedScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppliedScopeType(input string) (*AppliedScopeType, error) {
	vals := map[string]AppliedScopeType{
		"managementgroup": AppliedScopeTypeManagementGroup,
		"shared":          AppliedScopeTypeShared,
		"single":          AppliedScopeTypeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedScopeType(input)
	return &out, nil
}

type InstanceFlexibility string

const (
	InstanceFlexibilityOff InstanceFlexibility = "Off"
	InstanceFlexibilityOn  InstanceFlexibility = "On"
)

func PossibleValuesForInstanceFlexibility() []string {
	return []string{
		string(InstanceFlexibilityOff),
		string(InstanceFlexibilityOn),
	}
}

func (s *InstanceFlexibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstanceFlexibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstanceFlexibility(input string) (*InstanceFlexibility, error) {
	vals := map[string]InstanceFlexibility{
		"off": InstanceFlexibilityOff,
		"on":  InstanceFlexibilityOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstanceFlexibility(input)
	return &out, nil
}

type ReservationBillingPlan string

const (
	ReservationBillingPlanMonthly ReservationBillingPlan = "Monthly"
	ReservationBillingPlanUpfront ReservationBillingPlan = "Upfront"
)

func PossibleValuesForReservationBillingPlan() []string {
	return []string{
		string(ReservationBillingPlanMonthly),
		string(ReservationBillingPlanUpfront),
	}
}

func (s *ReservationBillingPlan) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationBillingPlan(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationBillingPlan(input string) (*ReservationBillingPlan, error) {
	vals := map[string]ReservationBillingPlan{
		"monthly": ReservationBillingPlanMonthly,
		"upfront": ReservationBillingPlanUpfront,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationBillingPlan(input)
	return &out, nil
}

type ReservationStatusCode string

const (
	ReservationStatusCodeActive                             ReservationStatusCode = "Active"
	ReservationStatusCodeCapacityError                      ReservationStatusCode = "CapacityError"
	ReservationStatusCodeCapacityRestricted                 ReservationStatusCode = "CapacityRestricted"
	ReservationStatusCodeCreditLineCheckFailed              ReservationStatusCode = "CreditLineCheckFailed"
	ReservationStatusCodeExchanged                          ReservationStatusCode = "Exchanged"
	ReservationStatusCodeExpired                            ReservationStatusCode = "Expired"
	ReservationStatusCodeMerged                             ReservationStatusCode = "Merged"
	ReservationStatusCodeNoBenefit                          ReservationStatusCode = "NoBenefit"
	ReservationStatusCodeNoBenefitDueToSubscriptionDeletion ReservationStatusCode = "NoBenefitDueToSubscriptionDeletion"
	ReservationStatusCodeNoBenefitDueToSubscriptionTransfer ReservationStatusCode = "NoBenefitDueToSubscriptionTransfer"
	ReservationStatusCodeNone                               ReservationStatusCode = "None"
	ReservationStatusCodePaymentInstrumentError             ReservationStatusCode = "PaymentInstrumentError"
	ReservationStatusCodePending                            ReservationStatusCode = "Pending"
	ReservationStatusCodeProcessing                         ReservationStatusCode = "Processing"
	ReservationStatusCodePurchaseError                      ReservationStatusCode = "PurchaseError"
	ReservationStatusCodeRiskCheckFailed                    ReservationStatusCode = "RiskCheckFailed"
	ReservationStatusCodeSplit                              ReservationStatusCode = "Split"
	ReservationStatusCodeSucceeded                          ReservationStatusCode = "Succeeded"
	ReservationStatusCodeUnknownError                       ReservationStatusCode = "UnknownError"
	ReservationStatusCodeWarning                            ReservationStatusCode = "Warning"
)

func PossibleValuesForReservationStatusCode() []string {
	return []string{
		string(ReservationStatusCodeActive),
		string(ReservationStatusCodeCapacityError),
		string(ReservationStatusCodeCapacityRestricted),
		string(ReservationStatusCodeCreditLineCheckFailed),
		string(ReservationStatusCodeExchanged),
		string(ReservationStatusCodeExpired),
		string(ReservationStatusCodeMerged),
		string(ReservationStatusCodeNoBenefit),
		string(ReservationStatusCodeNoBenefitDueToSubscriptionDeletion),
		string(ReservationStatusCodeNoBenefitDueToSubscriptionTransfer),
		string(ReservationStatusCodeNone),
		string(ReservationStatusCodePaymentInstrumentError),
		string(ReservationStatusCodePending),
		string(ReservationStatusCodeProcessing),
		string(ReservationStatusCodePurchaseError),
		string(ReservationStatusCodeRiskCheckFailed),
		string(ReservationStatusCodeSplit),
		string(ReservationStatusCodeSucceeded),
		string(ReservationStatusCodeUnknownError),
		string(ReservationStatusCodeWarning),
	}
}

func (s *ReservationStatusCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationStatusCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationStatusCode(input string) (*ReservationStatusCode, error) {
	vals := map[string]ReservationStatusCode{
		"active":                             ReservationStatusCodeActive,
		"capacityerror":                      ReservationStatusCodeCapacityError,
		"capacityrestricted":                 ReservationStatusCodeCapacityRestricted,
		"creditlinecheckfailed":              ReservationStatusCodeCreditLineCheckFailed,
		"exchanged":                          ReservationStatusCodeExchanged,
		"expired":                            ReservationStatusCodeExpired,
		"merged":                             ReservationStatusCodeMerged,
		"nobenefit":                          ReservationStatusCodeNoBenefit,
		"nobenefitduetosubscriptiondeletion": ReservationStatusCodeNoBenefitDueToSubscriptionDeletion,
		"nobenefitduetosubscriptiontransfer": ReservationStatusCodeNoBenefitDueToSubscriptionTransfer,
		"none":                               ReservationStatusCodeNone,
		"paymentinstrumenterror":             ReservationStatusCodePaymentInstrumentError,
		"pending":                            ReservationStatusCodePending,
		"processing":                         ReservationStatusCodeProcessing,
		"purchaseerror":                      ReservationStatusCodePurchaseError,
		"riskcheckfailed":                    ReservationStatusCodeRiskCheckFailed,
		"split":                              ReservationStatusCodeSplit,
		"succeeded":                          ReservationStatusCodeSucceeded,
		"unknownerror":                       ReservationStatusCodeUnknownError,
		"warning":                            ReservationStatusCodeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationStatusCode(input)
	return &out, nil
}
