package forecasts

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *Bound) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBound(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *ChargeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChargeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
