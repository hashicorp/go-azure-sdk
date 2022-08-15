package forecasts

import "strings"

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

func parseBound(input string) (*Bound, error) {
	vals := map[string]Bound{
		"lower": BoundLower,
		"upper": BoundUpper,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Bound(input)
	return &out, nil
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

func parseChargeType(input string) (*ChargeType, error) {
	vals := map[string]ChargeType{
		"actual":   ChargeTypeActual,
		"forecast": ChargeTypeForecast,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChargeType(input)
	return &out, nil
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

func parseGrain(input string) (*Grain, error) {
	vals := map[string]Grain{
		"daily":   GrainDaily,
		"monthly": GrainMonthly,
		"yearly":  GrainYearly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Grain(input)
	return &out, nil
}
