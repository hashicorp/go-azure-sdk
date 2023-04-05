package benefitrecommendations

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
