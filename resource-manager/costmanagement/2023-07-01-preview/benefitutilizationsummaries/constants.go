package benefitutilizationsummaries

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

func (s *GrainParameter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGrainParameter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGrainParameter(input string) (*GrainParameter, error) {
	vals := map[string]GrainParameter{
		"daily":   GrainParameterDaily,
		"hourly":  GrainParameterHourly,
		"monthly": GrainParameterMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GrainParameter(input)
	return &out, nil
}
