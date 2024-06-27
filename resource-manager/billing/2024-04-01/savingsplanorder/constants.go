package savingsplanorder

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPlan string

const (
	BillingPlanPOneM BillingPlan = "P1M"
)

func PossibleValuesForBillingPlan() []string {
	return []string{
		string(BillingPlanPOneM),
	}
}

func (s *BillingPlan) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingPlan(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingPlan(input string) (*BillingPlan, error) {
	vals := map[string]BillingPlan{
		"p1m": BillingPlanPOneM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingPlan(input)
	return &out, nil
}

type PaymentStatus string

const (
	PaymentStatusCancelled PaymentStatus = "Cancelled"
	PaymentStatusCompleted PaymentStatus = "Completed"
	PaymentStatusFailed    PaymentStatus = "Failed"
	PaymentStatusPending   PaymentStatus = "Pending"
	PaymentStatusScheduled PaymentStatus = "Scheduled"
	PaymentStatusSucceeded PaymentStatus = "Succeeded"
)

func PossibleValuesForPaymentStatus() []string {
	return []string{
		string(PaymentStatusCancelled),
		string(PaymentStatusCompleted),
		string(PaymentStatusFailed),
		string(PaymentStatusPending),
		string(PaymentStatusScheduled),
		string(PaymentStatusSucceeded),
	}
}

func (s *PaymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePaymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePaymentStatus(input string) (*PaymentStatus, error) {
	vals := map[string]PaymentStatus{
		"cancelled": PaymentStatusCancelled,
		"completed": PaymentStatusCompleted,
		"failed":    PaymentStatusFailed,
		"pending":   PaymentStatusPending,
		"scheduled": PaymentStatusScheduled,
		"succeeded": PaymentStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PaymentStatus(input)
	return &out, nil
}

type SavingsPlanTerm string

const (
	SavingsPlanTermPFiveY  SavingsPlanTerm = "P5Y"
	SavingsPlanTermPOneY   SavingsPlanTerm = "P1Y"
	SavingsPlanTermPThreeY SavingsPlanTerm = "P3Y"
)

func PossibleValuesForSavingsPlanTerm() []string {
	return []string{
		string(SavingsPlanTermPFiveY),
		string(SavingsPlanTermPOneY),
		string(SavingsPlanTermPThreeY),
	}
}

func (s *SavingsPlanTerm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSavingsPlanTerm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSavingsPlanTerm(input string) (*SavingsPlanTerm, error) {
	vals := map[string]SavingsPlanTerm{
		"p5y": SavingsPlanTermPFiveY,
		"p1y": SavingsPlanTermPOneY,
		"p3y": SavingsPlanTermPThreeY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SavingsPlanTerm(input)
	return &out, nil
}
