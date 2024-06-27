package savingsplan

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

type CommitmentGrain string

const (
	CommitmentGrainHourly CommitmentGrain = "Hourly"
)

func PossibleValuesForCommitmentGrain() []string {
	return []string{
		string(CommitmentGrainHourly),
	}
}

func (s *CommitmentGrain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommitmentGrain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommitmentGrain(input string) (*CommitmentGrain, error) {
	vals := map[string]CommitmentGrain{
		"hourly": CommitmentGrainHourly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommitmentGrain(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled         ProvisioningState = "Canceled"
	ProvisioningStateConfirmedBilling ProvisioningState = "ConfirmedBilling"
	ProvisioningStateCreated          ProvisioningState = "Created"
	ProvisioningStateCreating         ProvisioningState = "Creating"
	ProvisioningStateExpired          ProvisioningState = "Expired"
	ProvisioningStateFailed           ProvisioningState = "Failed"
	ProvisioningStatePendingBilling   ProvisioningState = "PendingBilling"
	ProvisioningStateSucceeded        ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateConfirmedBilling),
		string(ProvisioningStateCreated),
		string(ProvisioningStateCreating),
		string(ProvisioningStateExpired),
		string(ProvisioningStateFailed),
		string(ProvisioningStatePendingBilling),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":         ProvisioningStateCanceled,
		"confirmedbilling": ProvisioningStateConfirmedBilling,
		"created":          ProvisioningStateCreated,
		"creating":         ProvisioningStateCreating,
		"expired":          ProvisioningStateExpired,
		"failed":           ProvisioningStateFailed,
		"pendingbilling":   ProvisioningStatePendingBilling,
		"succeeded":        ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
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
