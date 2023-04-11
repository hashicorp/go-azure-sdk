package benefitrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitKind string

const (
	BenefitKindIncludedQuantity BenefitKind = "IncludedQuantity"
	BenefitKindReservation      BenefitKind = "Reservation"
	BenefitKindSavingsPlan      BenefitKind = "SavingsPlan"
)

func PossibleValuesForBenefitKind() []string {
	return []string{
		string(BenefitKindIncludedQuantity),
		string(BenefitKindReservation),
		string(BenefitKindSavingsPlan),
	}
}

func (s *BenefitKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBenefitKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBenefitKind(input string) (*BenefitKind, error) {
	vals := map[string]BenefitKind{
		"includedquantity": BenefitKindIncludedQuantity,
		"reservation":      BenefitKindReservation,
		"savingsplan":      BenefitKindSavingsPlan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BenefitKind(input)
	return &out, nil
}

type Grain string

const (
	GrainDaily   Grain = "Daily"
	GrainHourly  Grain = "Hourly"
	GrainMonthly Grain = "Monthly"
)

func PossibleValuesForGrain() []string {
	return []string{
		string(GrainDaily),
		string(GrainHourly),
		string(GrainMonthly),
	}
}

func (s *Grain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGrain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGrain(input string) (*Grain, error) {
	vals := map[string]Grain{
		"daily":   GrainDaily,
		"hourly":  GrainHourly,
		"monthly": GrainMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Grain(input)
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

type Term string

const (
	TermPOneY   Term = "P1Y"
	TermPThreeY Term = "P3Y"
)

func PossibleValuesForTerm() []string {
	return []string{
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
