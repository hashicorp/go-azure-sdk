package benefitutilizationsummaries

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

type GrainParameter string

const (
	GrainParameterDaily   GrainParameter = "Daily"
	GrainParameterHourly  GrainParameter = "Hourly"
	GrainParameterMonthly GrainParameter = "Monthly"
)

func PossibleValuesForGrainParameter() []string {
	return []string{
		string(GrainParameterDaily),
		string(GrainParameterHourly),
		string(GrainParameterMonthly),
	}
}
