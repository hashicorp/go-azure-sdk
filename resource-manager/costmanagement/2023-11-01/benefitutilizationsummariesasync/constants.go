package benefitutilizationsummariesasync

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

type BenefitUtilizationSummaryReportSchema string

const (
	BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage BenefitUtilizationSummaryReportSchema = "AvgUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaBenefitId                BenefitUtilizationSummaryReportSchema = "BenefitId"
	BenefitUtilizationSummaryReportSchemaBenefitOrderId           BenefitUtilizationSummaryReportSchema = "BenefitOrderId"
	BenefitUtilizationSummaryReportSchemaBenefitType              BenefitUtilizationSummaryReportSchema = "BenefitType"
	BenefitUtilizationSummaryReportSchemaKind                     BenefitUtilizationSummaryReportSchema = "Kind"
	BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage BenefitUtilizationSummaryReportSchema = "MaxUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage BenefitUtilizationSummaryReportSchema = "MinUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaUsageDate                BenefitUtilizationSummaryReportSchema = "UsageDate"
	BenefitUtilizationSummaryReportSchemaUtilizedPercentage       BenefitUtilizationSummaryReportSchema = "UtilizedPercentage"
)

func PossibleValuesForBenefitUtilizationSummaryReportSchema() []string {
	return []string{
		string(BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaBenefitId),
		string(BenefitUtilizationSummaryReportSchemaBenefitOrderId),
		string(BenefitUtilizationSummaryReportSchemaBenefitType),
		string(BenefitUtilizationSummaryReportSchemaKind),
		string(BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaUsageDate),
		string(BenefitUtilizationSummaryReportSchemaUtilizedPercentage),
	}
}

func (s *BenefitUtilizationSummaryReportSchema) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBenefitUtilizationSummaryReportSchema(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBenefitUtilizationSummaryReportSchema(input string) (*BenefitUtilizationSummaryReportSchema, error) {
	vals := map[string]BenefitUtilizationSummaryReportSchema{
		"avgutilizationpercentage": BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage,
		"benefitid":                BenefitUtilizationSummaryReportSchemaBenefitId,
		"benefitorderid":           BenefitUtilizationSummaryReportSchemaBenefitOrderId,
		"benefittype":              BenefitUtilizationSummaryReportSchemaBenefitType,
		"kind":                     BenefitUtilizationSummaryReportSchemaKind,
		"maxutilizationpercentage": BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage,
		"minutilizationpercentage": BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage,
		"usagedate":                BenefitUtilizationSummaryReportSchemaUsageDate,
		"utilizedpercentage":       BenefitUtilizationSummaryReportSchemaUtilizedPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BenefitUtilizationSummaryReportSchema(input)
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

type OperationStatusType string

const (
	OperationStatusTypeComplete OperationStatusType = "Complete"
	OperationStatusTypeFailed   OperationStatusType = "Failed"
	OperationStatusTypeRunning  OperationStatusType = "Running"
)

func PossibleValuesForOperationStatusType() []string {
	return []string{
		string(OperationStatusTypeComplete),
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
		"complete": OperationStatusTypeComplete,
		"failed":   OperationStatusTypeFailed,
		"running":  OperationStatusTypeRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusType(input)
	return &out, nil
}
