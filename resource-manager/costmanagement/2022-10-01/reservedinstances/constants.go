package reservedinstances

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
