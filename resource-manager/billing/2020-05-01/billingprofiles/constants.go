package billingprofiles

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileStatus string

const (
	BillingProfileStatusActive   BillingProfileStatus = "Active"
	BillingProfileStatusDisabled BillingProfileStatus = "Disabled"
	BillingProfileStatusWarned   BillingProfileStatus = "Warned"
)

func PossibleValuesForBillingProfileStatus() []string {
	return []string{
		string(BillingProfileStatusActive),
		string(BillingProfileStatusDisabled),
		string(BillingProfileStatusWarned),
	}
}

func (s *BillingProfileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatus(input string) (*BillingProfileStatus, error) {
	vals := map[string]BillingProfileStatus{
		"active":   BillingProfileStatusActive,
		"disabled": BillingProfileStatusDisabled,
		"warned":   BillingProfileStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatus(input)
	return &out, nil
}

type BillingRelationshipType string

const (
	BillingRelationshipTypeCSPPartner       BillingRelationshipType = "CSPPartner"
	BillingRelationshipTypeDirect           BillingRelationshipType = "Direct"
	BillingRelationshipTypeIndirectCustomer BillingRelationshipType = "IndirectCustomer"
	BillingRelationshipTypeIndirectPartner  BillingRelationshipType = "IndirectPartner"
)

func PossibleValuesForBillingRelationshipType() []string {
	return []string{
		string(BillingRelationshipTypeCSPPartner),
		string(BillingRelationshipTypeDirect),
		string(BillingRelationshipTypeIndirectCustomer),
		string(BillingRelationshipTypeIndirectPartner),
	}
}

func (s *BillingRelationshipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingRelationshipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingRelationshipType(input string) (*BillingRelationshipType, error) {
	vals := map[string]BillingRelationshipType{
		"csppartner":       BillingRelationshipTypeCSPPartner,
		"direct":           BillingRelationshipTypeDirect,
		"indirectcustomer": BillingRelationshipTypeIndirectCustomer,
		"indirectpartner":  BillingRelationshipTypeIndirectPartner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRelationshipType(input)
	return &out, nil
}

type InvoiceSectionState string

const (
	InvoiceSectionStateActive     InvoiceSectionState = "Active"
	InvoiceSectionStateRestricted InvoiceSectionState = "Restricted"
)

func PossibleValuesForInvoiceSectionState() []string {
	return []string{
		string(InvoiceSectionStateActive),
		string(InvoiceSectionStateRestricted),
	}
}

func (s *InvoiceSectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceSectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceSectionState(input string) (*InvoiceSectionState, error) {
	vals := map[string]InvoiceSectionState{
		"active":     InvoiceSectionStateActive,
		"restricted": InvoiceSectionStateRestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionState(input)
	return &out, nil
}

type SpendingLimit string

const (
	SpendingLimitOff SpendingLimit = "Off"
	SpendingLimitOn  SpendingLimit = "On"
)

func PossibleValuesForSpendingLimit() []string {
	return []string{
		string(SpendingLimitOff),
		string(SpendingLimitOn),
	}
}

func (s *SpendingLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimit(input string) (*SpendingLimit, error) {
	vals := map[string]SpendingLimit{
		"off": SpendingLimitOff,
		"on":  SpendingLimitOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimit(input)
	return &out, nil
}

type StatusReasonCode string

const (
	StatusReasonCodePastDue              StatusReasonCode = "PastDue"
	StatusReasonCodeSpendingLimitExpired StatusReasonCode = "SpendingLimitExpired"
	StatusReasonCodeSpendingLimitReached StatusReasonCode = "SpendingLimitReached"
)

func PossibleValuesForStatusReasonCode() []string {
	return []string{
		string(StatusReasonCodePastDue),
		string(StatusReasonCodeSpendingLimitExpired),
		string(StatusReasonCodeSpendingLimitReached),
	}
}

func (s *StatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusReasonCode(input string) (*StatusReasonCode, error) {
	vals := map[string]StatusReasonCode{
		"pastdue":              StatusReasonCodePastDue,
		"spendinglimitexpired": StatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": StatusReasonCodeSpendingLimitReached,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusReasonCode(input)
	return &out, nil
}

type TargetCloud string

const (
	TargetCloudUSGov TargetCloud = "USGov"
	TargetCloudUSNat TargetCloud = "USNat"
	TargetCloudUSSec TargetCloud = "USSec"
)

func PossibleValuesForTargetCloud() []string {
	return []string{
		string(TargetCloudUSGov),
		string(TargetCloudUSNat),
		string(TargetCloudUSSec),
	}
}

func (s *TargetCloud) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetCloud(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetCloud(input string) (*TargetCloud, error) {
	vals := map[string]TargetCloud{
		"usgov": TargetCloudUSGov,
		"usnat": TargetCloudUSNat,
		"ussec": TargetCloudUSSec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetCloud(input)
	return &out, nil
}
