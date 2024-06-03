package reservedinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

func PossibleValuesForOperationStatusType() []string {
	return []string{
		string(OperationStatusTypeCompleted),
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
		"completed": OperationStatusTypeCompleted,
		"failed":    OperationStatusTypeFailed,
		"running":   OperationStatusTypeRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusType(input)
	return &out, nil
}

type ReservationReportSchema string

const (
	ReservationReportSchemaInstanceFlexibilityGroup ReservationReportSchema = "InstanceFlexibilityGroup"
	ReservationReportSchemaInstanceFlexibilityRatio ReservationReportSchema = "InstanceFlexibilityRatio"
	ReservationReportSchemaInstanceId               ReservationReportSchema = "InstanceId"
	ReservationReportSchemaKind                     ReservationReportSchema = "Kind"
	ReservationReportSchemaReservationId            ReservationReportSchema = "ReservationId"
	ReservationReportSchemaReservationOrderId       ReservationReportSchema = "ReservationOrderId"
	ReservationReportSchemaReservedHours            ReservationReportSchema = "ReservedHours"
	ReservationReportSchemaSkuName                  ReservationReportSchema = "SkuName"
	ReservationReportSchemaTotalReservedQuantity    ReservationReportSchema = "TotalReservedQuantity"
	ReservationReportSchemaUsageDate                ReservationReportSchema = "UsageDate"
	ReservationReportSchemaUsedHours                ReservationReportSchema = "UsedHours"
)

func PossibleValuesForReservationReportSchema() []string {
	return []string{
		string(ReservationReportSchemaInstanceFlexibilityGroup),
		string(ReservationReportSchemaInstanceFlexibilityRatio),
		string(ReservationReportSchemaInstanceId),
		string(ReservationReportSchemaKind),
		string(ReservationReportSchemaReservationId),
		string(ReservationReportSchemaReservationOrderId),
		string(ReservationReportSchemaReservedHours),
		string(ReservationReportSchemaSkuName),
		string(ReservationReportSchemaTotalReservedQuantity),
		string(ReservationReportSchemaUsageDate),
		string(ReservationReportSchemaUsedHours),
	}
}

func (s *ReservationReportSchema) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationReportSchema(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationReportSchema(input string) (*ReservationReportSchema, error) {
	vals := map[string]ReservationReportSchema{
		"instanceflexibilitygroup": ReservationReportSchemaInstanceFlexibilityGroup,
		"instanceflexibilityratio": ReservationReportSchemaInstanceFlexibilityRatio,
		"instanceid":               ReservationReportSchemaInstanceId,
		"kind":                     ReservationReportSchemaKind,
		"reservationid":            ReservationReportSchemaReservationId,
		"reservationorderid":       ReservationReportSchemaReservationOrderId,
		"reservedhours":            ReservationReportSchemaReservedHours,
		"skuname":                  ReservationReportSchemaSkuName,
		"totalreservedquantity":    ReservationReportSchemaTotalReservedQuantity,
		"usagedate":                ReservationReportSchemaUsageDate,
		"usedhours":                ReservationReportSchemaUsedHours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationReportSchema(input)
	return &out, nil
}
