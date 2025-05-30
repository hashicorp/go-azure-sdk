package exports

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompressionModeType string

const (
	CompressionModeTypeGzip   CompressionModeType = "gzip"
	CompressionModeTypeNone   CompressionModeType = "none"
	CompressionModeTypeSnappy CompressionModeType = "snappy"
)

func PossibleValuesForCompressionModeType() []string {
	return []string{
		string(CompressionModeTypeGzip),
		string(CompressionModeTypeNone),
		string(CompressionModeTypeSnappy),
	}
}

func (s *CompressionModeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompressionModeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompressionModeType(input string) (*CompressionModeType, error) {
	vals := map[string]CompressionModeType{
		"gzip":   CompressionModeTypeGzip,
		"none":   CompressionModeTypeNone,
		"snappy": CompressionModeTypeSnappy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompressionModeType(input)
	return &out, nil
}

type DataOverwriteBehaviorType string

const (
	DataOverwriteBehaviorTypeCreateNewReport         DataOverwriteBehaviorType = "CreateNewReport"
	DataOverwriteBehaviorTypeOverwritePreviousReport DataOverwriteBehaviorType = "OverwritePreviousReport"
)

func PossibleValuesForDataOverwriteBehaviorType() []string {
	return []string{
		string(DataOverwriteBehaviorTypeCreateNewReport),
		string(DataOverwriteBehaviorTypeOverwritePreviousReport),
	}
}

func (s *DataOverwriteBehaviorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataOverwriteBehaviorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataOverwriteBehaviorType(input string) (*DataOverwriteBehaviorType, error) {
	vals := map[string]DataOverwriteBehaviorType{
		"createnewreport":         DataOverwriteBehaviorTypeCreateNewReport,
		"overwritepreviousreport": DataOverwriteBehaviorTypeOverwritePreviousReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataOverwriteBehaviorType(input)
	return &out, nil
}

type DestinationType string

const (
	DestinationTypeAzureBlob DestinationType = "AzureBlob"
)

func PossibleValuesForDestinationType() []string {
	return []string{
		string(DestinationTypeAzureBlob),
	}
}

func (s *DestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDestinationType(input string) (*DestinationType, error) {
	vals := map[string]DestinationType{
		"azureblob": DestinationTypeAzureBlob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DestinationType(input)
	return &out, nil
}

type ExecutionStatus string

const (
	ExecutionStatusCompleted           ExecutionStatus = "Completed"
	ExecutionStatusDataNotAvailable    ExecutionStatus = "DataNotAvailable"
	ExecutionStatusFailed              ExecutionStatus = "Failed"
	ExecutionStatusInProgress          ExecutionStatus = "InProgress"
	ExecutionStatusNewDataNotAvailable ExecutionStatus = "NewDataNotAvailable"
	ExecutionStatusQueued              ExecutionStatus = "Queued"
	ExecutionStatusTimeout             ExecutionStatus = "Timeout"
)

func PossibleValuesForExecutionStatus() []string {
	return []string{
		string(ExecutionStatusCompleted),
		string(ExecutionStatusDataNotAvailable),
		string(ExecutionStatusFailed),
		string(ExecutionStatusInProgress),
		string(ExecutionStatusNewDataNotAvailable),
		string(ExecutionStatusQueued),
		string(ExecutionStatusTimeout),
	}
}

func (s *ExecutionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExecutionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExecutionStatus(input string) (*ExecutionStatus, error) {
	vals := map[string]ExecutionStatus{
		"completed":           ExecutionStatusCompleted,
		"datanotavailable":    ExecutionStatusDataNotAvailable,
		"failed":              ExecutionStatusFailed,
		"inprogress":          ExecutionStatusInProgress,
		"newdatanotavailable": ExecutionStatusNewDataNotAvailable,
		"queued":              ExecutionStatusQueued,
		"timeout":             ExecutionStatusTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionStatus(input)
	return &out, nil
}

type ExecutionType string

const (
	ExecutionTypeOnDemand  ExecutionType = "OnDemand"
	ExecutionTypeScheduled ExecutionType = "Scheduled"
)

func PossibleValuesForExecutionType() []string {
	return []string{
		string(ExecutionTypeOnDemand),
		string(ExecutionTypeScheduled),
	}
}

func (s *ExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExecutionType(input string) (*ExecutionType, error) {
	vals := map[string]ExecutionType{
		"ondemand":  ExecutionTypeOnDemand,
		"scheduled": ExecutionTypeScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionType(input)
	return &out, nil
}

type ExportType string

const (
	ExportTypeActualCost                 ExportType = "ActualCost"
	ExportTypeAmortizedCost              ExportType = "AmortizedCost"
	ExportTypeFocusCost                  ExportType = "FocusCost"
	ExportTypePriceSheet                 ExportType = "PriceSheet"
	ExportTypeReservationDetails         ExportType = "ReservationDetails"
	ExportTypeReservationRecommendations ExportType = "ReservationRecommendations"
	ExportTypeReservationTransactions    ExportType = "ReservationTransactions"
	ExportTypeUsage                      ExportType = "Usage"
)

func PossibleValuesForExportType() []string {
	return []string{
		string(ExportTypeActualCost),
		string(ExportTypeAmortizedCost),
		string(ExportTypeFocusCost),
		string(ExportTypePriceSheet),
		string(ExportTypeReservationDetails),
		string(ExportTypeReservationRecommendations),
		string(ExportTypeReservationTransactions),
		string(ExportTypeUsage),
	}
}

func (s *ExportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExportType(input string) (*ExportType, error) {
	vals := map[string]ExportType{
		"actualcost":                 ExportTypeActualCost,
		"amortizedcost":              ExportTypeAmortizedCost,
		"focuscost":                  ExportTypeFocusCost,
		"pricesheet":                 ExportTypePriceSheet,
		"reservationdetails":         ExportTypeReservationDetails,
		"reservationrecommendations": ExportTypeReservationRecommendations,
		"reservationtransactions":    ExportTypeReservationTransactions,
		"usage":                      ExportTypeUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportType(input)
	return &out, nil
}

type FilterItemNames string

const (
	FilterItemNamesLookBackPeriod   FilterItemNames = "LookBackPeriod"
	FilterItemNamesReservationScope FilterItemNames = "ReservationScope"
	FilterItemNamesResourceType     FilterItemNames = "ResourceType"
)

func PossibleValuesForFilterItemNames() []string {
	return []string{
		string(FilterItemNamesLookBackPeriod),
		string(FilterItemNamesReservationScope),
		string(FilterItemNamesResourceType),
	}
}

func (s *FilterItemNames) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterItemNames(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterItemNames(input string) (*FilterItemNames, error) {
	vals := map[string]FilterItemNames{
		"lookbackperiod":   FilterItemNamesLookBackPeriod,
		"reservationscope": FilterItemNamesReservationScope,
		"resourcetype":     FilterItemNamesResourceType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterItemNames(input)
	return &out, nil
}

type FormatType string

const (
	FormatTypeCsv     FormatType = "Csv"
	FormatTypeParquet FormatType = "Parquet"
)

func PossibleValuesForFormatType() []string {
	return []string{
		string(FormatTypeCsv),
		string(FormatTypeParquet),
	}
}

func (s *FormatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFormatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFormatType(input string) (*FormatType, error) {
	vals := map[string]FormatType{
		"csv":     FormatTypeCsv,
		"parquet": FormatTypeParquet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FormatType(input)
	return &out, nil
}

type GranularityType string

const (
	GranularityTypeDaily   GranularityType = "Daily"
	GranularityTypeMonthly GranularityType = "Monthly"
)

func PossibleValuesForGranularityType() []string {
	return []string{
		string(GranularityTypeDaily),
		string(GranularityTypeMonthly),
	}
}

func (s *GranularityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGranularityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGranularityType(input string) (*GranularityType, error) {
	vals := map[string]GranularityType{
		"daily":   GranularityTypeDaily,
		"monthly": GranularityTypeMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GranularityType(input)
	return &out, nil
}

type RecurrenceType string

const (
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	RecurrenceTypeDaily    RecurrenceType = "Daily"
	RecurrenceTypeMonthly  RecurrenceType = "Monthly"
	RecurrenceTypeWeekly   RecurrenceType = "Weekly"
)

func PossibleValuesForRecurrenceType() []string {
	return []string{
		string(RecurrenceTypeAnnually),
		string(RecurrenceTypeDaily),
		string(RecurrenceTypeMonthly),
		string(RecurrenceTypeWeekly),
	}
}

func (s *RecurrenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrenceType(input string) (*RecurrenceType, error) {
	vals := map[string]RecurrenceType{
		"annually": RecurrenceTypeAnnually,
		"daily":    RecurrenceTypeDaily,
		"monthly":  RecurrenceTypeMonthly,
		"weekly":   RecurrenceTypeWeekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceType(input)
	return &out, nil
}

type StatusType string

const (
	StatusTypeActive   StatusType = "Active"
	StatusTypeInactive StatusType = "Inactive"
)

func PossibleValuesForStatusType() []string {
	return []string{
		string(StatusTypeActive),
		string(StatusTypeInactive),
	}
}

func (s *StatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusType(input string) (*StatusType, error) {
	vals := map[string]StatusType{
		"active":   StatusTypeActive,
		"inactive": StatusTypeInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusType(input)
	return &out, nil
}

type TimeframeType string

const (
	TimeframeTypeBillingMonthToDate  TimeframeType = "BillingMonthToDate"
	TimeframeTypeCustom              TimeframeType = "Custom"
	TimeframeTypeMonthToDate         TimeframeType = "MonthToDate"
	TimeframeTypeTheCurrentMonth     TimeframeType = "TheCurrentMonth"
	TimeframeTypeTheLastBillingMonth TimeframeType = "TheLastBillingMonth"
	TimeframeTypeTheLastMonth        TimeframeType = "TheLastMonth"
	TimeframeTypeWeekToDate          TimeframeType = "WeekToDate"
)

func PossibleValuesForTimeframeType() []string {
	return []string{
		string(TimeframeTypeBillingMonthToDate),
		string(TimeframeTypeCustom),
		string(TimeframeTypeMonthToDate),
		string(TimeframeTypeTheCurrentMonth),
		string(TimeframeTypeTheLastBillingMonth),
		string(TimeframeTypeTheLastMonth),
		string(TimeframeTypeWeekToDate),
	}
}

func (s *TimeframeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeframeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeframeType(input string) (*TimeframeType, error) {
	vals := map[string]TimeframeType{
		"billingmonthtodate":  TimeframeTypeBillingMonthToDate,
		"custom":              TimeframeTypeCustom,
		"monthtodate":         TimeframeTypeMonthToDate,
		"thecurrentmonth":     TimeframeTypeTheCurrentMonth,
		"thelastbillingmonth": TimeframeTypeTheLastBillingMonth,
		"thelastmonth":        TimeframeTypeTheLastMonth,
		"weektodate":          TimeframeTypeWeekToDate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeframeType(input)
	return &out, nil
}
