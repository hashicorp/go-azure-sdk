package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingFrequency string

const (
	BillingFrequencyMonth   BillingFrequency = "Month"
	BillingFrequencyQuarter BillingFrequency = "Quarter"
	BillingFrequencyYear    BillingFrequency = "Year"
)

func PossibleValuesForBillingFrequency() []string {
	return []string{
		string(BillingFrequencyMonth),
		string(BillingFrequencyQuarter),
		string(BillingFrequencyYear),
	}
}

func (s *BillingFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingFrequency(input string) (*BillingFrequency, error) {
	vals := map[string]BillingFrequency{
		"month":   BillingFrequencyMonth,
		"quarter": BillingFrequencyQuarter,
		"year":    BillingFrequencyYear,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingFrequency(input)
	return &out, nil
}

type ChargeSummaryKind string

const (
	ChargeSummaryKindLegacy ChargeSummaryKind = "legacy"
	ChargeSummaryKindModern ChargeSummaryKind = "modern"
)

func PossibleValuesForChargeSummaryKind() []string {
	return []string{
		string(ChargeSummaryKindLegacy),
		string(ChargeSummaryKindModern),
	}
}

func (s *ChargeSummaryKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChargeSummaryKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChargeSummaryKind(input string) (*ChargeSummaryKind, error) {
	vals := map[string]ChargeSummaryKind{
		"legacy": ChargeSummaryKindLegacy,
		"modern": ChargeSummaryKindModern,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChargeSummaryKind(input)
	return &out, nil
}

type Datagrain string

const (
	DatagrainDaily   Datagrain = "daily"
	DatagrainMonthly Datagrain = "monthly"
)

func PossibleValuesForDatagrain() []string {
	return []string{
		string(DatagrainDaily),
		string(DatagrainMonthly),
	}
}

func (s *Datagrain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatagrain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatagrain(input string) (*Datagrain, error) {
	vals := map[string]Datagrain{
		"daily":   DatagrainDaily,
		"monthly": DatagrainMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Datagrain(input)
	return &out, nil
}

type EventType string

const (
	EventTypeCreditExpired        EventType = "CreditExpired"
	EventTypeNewCredit            EventType = "NewCredit"
	EventTypePendingAdjustments   EventType = "PendingAdjustments"
	EventTypePendingCharges       EventType = "PendingCharges"
	EventTypePendingExpiredCredit EventType = "PendingExpiredCredit"
	EventTypePendingNewCredit     EventType = "PendingNewCredit"
	EventTypeSettledCharges       EventType = "SettledCharges"
	EventTypeUnKnown              EventType = "UnKnown"
)

func PossibleValuesForEventType() []string {
	return []string{
		string(EventTypeCreditExpired),
		string(EventTypeNewCredit),
		string(EventTypePendingAdjustments),
		string(EventTypePendingCharges),
		string(EventTypePendingExpiredCredit),
		string(EventTypePendingNewCredit),
		string(EventTypeSettledCharges),
		string(EventTypeUnKnown),
	}
}

func (s *EventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"creditexpired":        EventTypeCreditExpired,
		"newcredit":            EventTypeNewCredit,
		"pendingadjustments":   EventTypePendingAdjustments,
		"pendingcharges":       EventTypePendingCharges,
		"pendingexpiredcredit": EventTypePendingExpiredCredit,
		"pendingnewcredit":     EventTypePendingNewCredit,
		"settledcharges":       EventTypeSettledCharges,
		"unknown":              EventTypeUnKnown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventType(input)
	return &out, nil
}

type LookBackPeriod string

const (
	LookBackPeriodLastSevenDays     LookBackPeriod = "Last7Days"
	LookBackPeriodLastSixZeroDays   LookBackPeriod = "Last60Days"
	LookBackPeriodLastThreeZeroDays LookBackPeriod = "Last30Days"
)

func PossibleValuesForLookBackPeriod() []string {
	return []string{
		string(LookBackPeriodLastSevenDays),
		string(LookBackPeriodLastSixZeroDays),
		string(LookBackPeriodLastThreeZeroDays),
	}
}

func (s *LookBackPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLookBackPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLookBackPeriod(input string) (*LookBackPeriod, error) {
	vals := map[string]LookBackPeriod{
		"last7days":  LookBackPeriodLastSevenDays,
		"last60days": LookBackPeriodLastSixZeroDays,
		"last30days": LookBackPeriodLastThreeZeroDays,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LookBackPeriod(input)
	return &out, nil
}

type LotSource string

const (
	LotSourceConsumptionCommitment LotSource = "ConsumptionCommitment"
	LotSourcePromotionalCredit     LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit       LotSource = "PurchasedCredit"
)

func PossibleValuesForLotSource() []string {
	return []string{
		string(LotSourceConsumptionCommitment),
		string(LotSourcePromotionalCredit),
		string(LotSourcePurchasedCredit),
	}
}

func (s *LotSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLotSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLotSource(input string) (*LotSource, error) {
	vals := map[string]LotSource{
		"consumptioncommitment": LotSourceConsumptionCommitment,
		"promotionalcredit":     LotSourcePromotionalCredit,
		"purchasedcredit":       LotSourcePurchasedCredit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LotSource(input)
	return &out, nil
}

type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

func PossibleValuesForOperationStatusType() []string {
	return []string{
		string(OperationStatusTypeCompleted),
		string(OperationStatusTypeFailed),
		string(OperationStatusTypeRunning),
	}
}

func (s *OperationStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationStatusType(input string) (*OperationStatusType, error) {
	vals := map[string]OperationStatusType{
		"completed": OperationStatusTypeCompleted,
		"failed":    OperationStatusTypeFailed,
		"running":   OperationStatusTypeRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusType(input)
	return &out, nil
}

type OrganizationType string

const (
	OrganizationTypeContributor OrganizationType = "Contributor"
	OrganizationTypePrimary     OrganizationType = "Primary"
)

func PossibleValuesForOrganizationType() []string {
	return []string{
		string(OrganizationTypeContributor),
		string(OrganizationTypePrimary),
	}
}

func (s *OrganizationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrganizationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrganizationType(input string) (*OrganizationType, error) {
	vals := map[string]OrganizationType{
		"contributor": OrganizationTypeContributor,
		"primary":     OrganizationTypePrimary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationType(input)
	return &out, nil
}

type ReservationRecommendationKind string

const (
	ReservationRecommendationKindLegacy ReservationRecommendationKind = "legacy"
	ReservationRecommendationKindModern ReservationRecommendationKind = "modern"
)

func PossibleValuesForReservationRecommendationKind() []string {
	return []string{
		string(ReservationRecommendationKindLegacy),
		string(ReservationRecommendationKindModern),
	}
}

func (s *ReservationRecommendationKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationRecommendationKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationRecommendationKind(input string) (*ReservationRecommendationKind, error) {
	vals := map[string]ReservationRecommendationKind{
		"legacy": ReservationRecommendationKindLegacy,
		"modern": ReservationRecommendationKindModern,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationRecommendationKind(input)
	return &out, nil
}

type Scope string

const (
	ScopeShared Scope = "Shared"
	ScopeSingle Scope = "Single"
)

func PossibleValuesForScope() []string {
	return []string{
		string(ScopeShared),
		string(ScopeSingle),
	}
}

func (s *Scope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScope(input string) (*Scope, error) {
	vals := map[string]Scope{
		"shared": ScopeShared,
		"single": ScopeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Scope(input)
	return &out, nil
}

type Status string

const (
	StatusActive   Status = "Active"
	StatusCanceled Status = "Canceled"
	StatusComplete Status = "Complete"
	StatusExpired  Status = "Expired"
	StatusInactive Status = "Inactive"
	StatusNone     Status = "None"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusActive),
		string(StatusCanceled),
		string(StatusComplete),
		string(StatusExpired),
		string(StatusInactive),
		string(StatusNone),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"active":   StatusActive,
		"canceled": StatusCanceled,
		"complete": StatusComplete,
		"expired":  StatusExpired,
		"inactive": StatusInactive,
		"none":     StatusNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type Term string

const (
	TermPOneM   Term = "P1M"
	TermPOneY   Term = "P1Y"
	TermPThreeY Term = "P3Y"
)

func PossibleValuesForTerm() []string {
	return []string{
		string(TermPOneM),
		string(TermPOneY),
		string(TermPThreeY),
	}
}

func (s *Term) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTerm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTerm(input string) (*Term, error) {
	vals := map[string]Term{
		"p1m": TermPOneM,
		"p1y": TermPOneY,
		"p3y": TermPThreeY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Term(input)
	return &out, nil
}
