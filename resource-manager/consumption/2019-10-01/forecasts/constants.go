package forecasts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Bound string

const (
	BoundLower Bound = "Lower"
	BoundUpper Bound = "Upper"
)

func PossibleValuesForBound() []string {
	return []string{
		string(BoundLower),
		string(BoundUpper),
	}
}

type ChargeType string

const (
	ChargeTypeActual   ChargeType = "Actual"
	ChargeTypeForecast ChargeType = "Forecast"
)

func PossibleValuesForChargeType() []string {
	return []string{
		string(ChargeTypeActual),
		string(ChargeTypeForecast),
	}
}

type Grain string

const (
	GrainDaily   Grain = "Daily"
	GrainMonthly Grain = "Monthly"
	GrainYearly  Grain = "Yearly"
)

func PossibleValuesForGrain() []string {
	return []string{
		string(GrainDaily),
		string(GrainMonthly),
		string(GrainYearly),
	}
}
