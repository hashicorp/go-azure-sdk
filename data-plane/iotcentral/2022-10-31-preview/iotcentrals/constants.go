package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityAggregateType string

const (
	CapabilityAggregateTypeAvg   CapabilityAggregateType = "avg"
	CapabilityAggregateTypeCount CapabilityAggregateType = "count"
	CapabilityAggregateTypeMax   CapabilityAggregateType = "max"
	CapabilityAggregateTypeMin   CapabilityAggregateType = "min"
	CapabilityAggregateTypeSum   CapabilityAggregateType = "sum"
)

func PossibleValuesForCapabilityAggregateType() []string {
	return []string{
		string(CapabilityAggregateTypeAvg),
		string(CapabilityAggregateTypeCount),
		string(CapabilityAggregateTypeMax),
		string(CapabilityAggregateTypeMin),
		string(CapabilityAggregateTypeSum),
	}
}

func (s *CapabilityAggregateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCapabilityAggregateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCapabilityAggregateType(input string) (*CapabilityAggregateType, error) {
	vals := map[string]CapabilityAggregateType{
		"avg":   CapabilityAggregateTypeAvg,
		"count": CapabilityAggregateTypeCount,
		"max":   CapabilityAggregateTypeMax,
		"min":   CapabilityAggregateTypeMin,
		"sum":   CapabilityAggregateTypeSum,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CapabilityAggregateType(input)
	return &out, nil
}

type CertificateEntry string

const (
	CertificateEntryPrimary   CertificateEntry = "primary"
	CertificateEntrySecondary CertificateEntry = "secondary"
)

func PossibleValuesForCertificateEntry() []string {
	return []string{
		string(CertificateEntryPrimary),
		string(CertificateEntrySecondary),
	}
}

func (s *CertificateEntry) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateEntry(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateEntry(input string) (*CertificateEntry, error) {
	vals := map[string]CertificateEntry{
		"primary":   CertificateEntryPrimary,
		"secondary": CertificateEntrySecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateEntry(input)
	return &out, nil
}

type DestinationSource string

const (
	DestinationSourceAudit                   DestinationSource = "audit"
	DestinationSourceDeviceConnectivity      DestinationSource = "deviceConnectivity"
	DestinationSourceDeviceLifecycle         DestinationSource = "deviceLifecycle"
	DestinationSourceDeviceTemplateLifecycle DestinationSource = "deviceTemplateLifecycle"
	DestinationSourceProperties              DestinationSource = "properties"
	DestinationSourceTelemetry               DestinationSource = "telemetry"
)

func PossibleValuesForDestinationSource() []string {
	return []string{
		string(DestinationSourceAudit),
		string(DestinationSourceDeviceConnectivity),
		string(DestinationSourceDeviceLifecycle),
		string(DestinationSourceDeviceTemplateLifecycle),
		string(DestinationSourceProperties),
		string(DestinationSourceTelemetry),
	}
}

func (s *DestinationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDestinationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDestinationSource(input string) (*DestinationSource, error) {
	vals := map[string]DestinationSource{
		"audit":                   DestinationSourceAudit,
		"deviceconnectivity":      DestinationSourceDeviceConnectivity,
		"devicelifecycle":         DestinationSourceDeviceLifecycle,
		"devicetemplatelifecycle": DestinationSourceDeviceTemplateLifecycle,
		"properties":              DestinationSourceProperties,
		"telemetry":               DestinationSourceTelemetry,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DestinationSource(input)
	return &out, nil
}

type DeviceType string

const (
	DeviceTypeIotEdge DeviceType = "iotEdge"
)

func PossibleValuesForDeviceType() []string {
	return []string{
		string(DeviceTypeIotEdge),
	}
}

func (s *DeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceType(input string) (*DeviceType, error) {
	vals := map[string]DeviceType{
		"iotedge": DeviceTypeIotEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceType(input)
	return &out, nil
}

type EnrollmentGroupType string

const (
	EnrollmentGroupTypeIot     EnrollmentGroupType = "iot"
	EnrollmentGroupTypeIotEdge EnrollmentGroupType = "iotEdge"
)

func PossibleValuesForEnrollmentGroupType() []string {
	return []string{
		string(EnrollmentGroupTypeIot),
		string(EnrollmentGroupTypeIotEdge),
	}
}

func (s *EnrollmentGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentGroupType(input string) (*EnrollmentGroupType, error) {
	vals := map[string]EnrollmentGroupType{
		"iot":     EnrollmentGroupTypeIot,
		"iotedge": EnrollmentGroupTypeIotEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentGroupType(input)
	return &out, nil
}

type FileUploadState string

const (
	FileUploadStateDeleting  FileUploadState = "deleting"
	FileUploadStateFailed    FileUploadState = "failed"
	FileUploadStatePending   FileUploadState = "pending"
	FileUploadStateSucceeded FileUploadState = "succeeded"
	FileUploadStateUpdating  FileUploadState = "updating"
)

func PossibleValuesForFileUploadState() []string {
	return []string{
		string(FileUploadStateDeleting),
		string(FileUploadStateFailed),
		string(FileUploadStatePending),
		string(FileUploadStateSucceeded),
		string(FileUploadStateUpdating),
	}
}

func (s *FileUploadState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileUploadState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileUploadState(input string) (*FileUploadState, error) {
	vals := map[string]FileUploadState{
		"deleting":  FileUploadStateDeleting,
		"failed":    FileUploadStateFailed,
		"pending":   FileUploadStatePending,
		"succeeded": FileUploadStateSucceeded,
		"updating":  FileUploadStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileUploadState(input)
	return &out, nil
}

type ImageTileTextUnits string

const (
	ImageTileTextUnitsPx ImageTileTextUnits = "px"
)

func PossibleValuesForImageTileTextUnits() []string {
	return []string{
		string(ImageTileTextUnitsPx),
	}
}

func (s *ImageTileTextUnits) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImageTileTextUnits(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImageTileTextUnits(input string) (*ImageTileTextUnits, error) {
	vals := map[string]ImageTileTextUnits{
		"px": ImageTileTextUnitsPx,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImageTileTextUnits(input)
	return &out, nil
}

type JobBatchType string

const (
	JobBatchTypeNumber     JobBatchType = "number"
	JobBatchTypePercentage JobBatchType = "percentage"
)

func PossibleValuesForJobBatchType() []string {
	return []string{
		string(JobBatchTypeNumber),
		string(JobBatchTypePercentage),
	}
}

func (s *JobBatchType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobBatchType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobBatchType(input string) (*JobBatchType, error) {
	vals := map[string]JobBatchType{
		"number":     JobBatchTypeNumber,
		"percentage": JobBatchTypePercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobBatchType(input)
	return &out, nil
}

type JobCancellationThresholdType string

const (
	JobCancellationThresholdTypeNumber     JobCancellationThresholdType = "number"
	JobCancellationThresholdTypePercentage JobCancellationThresholdType = "percentage"
)

func PossibleValuesForJobCancellationThresholdType() []string {
	return []string{
		string(JobCancellationThresholdTypeNumber),
		string(JobCancellationThresholdTypePercentage),
	}
}

func (s *JobCancellationThresholdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobCancellationThresholdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobCancellationThresholdType(input string) (*JobCancellationThresholdType, error) {
	vals := map[string]JobCancellationThresholdType{
		"number":     JobCancellationThresholdTypeNumber,
		"percentage": JobCancellationThresholdTypePercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobCancellationThresholdType(input)
	return &out, nil
}

type JobRecurrence string

const (
	JobRecurrenceDaily   JobRecurrence = "daily"
	JobRecurrenceMonthly JobRecurrence = "monthly"
	JobRecurrenceWeekly  JobRecurrence = "weekly"
)

func PossibleValuesForJobRecurrence() []string {
	return []string{
		string(JobRecurrenceDaily),
		string(JobRecurrenceMonthly),
		string(JobRecurrenceWeekly),
	}
}

func (s *JobRecurrence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobRecurrence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobRecurrence(input string) (*JobRecurrence, error) {
	vals := map[string]JobRecurrence{
		"daily":   JobRecurrenceDaily,
		"monthly": JobRecurrenceMonthly,
		"weekly":  JobRecurrenceWeekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobRecurrence(input)
	return &out, nil
}

type OAuthRequestType string

const (
	OAuthRequestTypeAuto       OAuthRequestType = "auto"
	OAuthRequestTypeJson       OAuthRequestType = "json"
	OAuthRequestTypeUrlencoded OAuthRequestType = "urlencoded"
)

func PossibleValuesForOAuthRequestType() []string {
	return []string{
		string(OAuthRequestTypeAuto),
		string(OAuthRequestTypeJson),
		string(OAuthRequestTypeUrlencoded),
	}
}

func (s *OAuthRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOAuthRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOAuthRequestType(input string) (*OAuthRequestType, error) {
	vals := map[string]OAuthRequestType{
		"auto":       OAuthRequestTypeAuto,
		"json":       OAuthRequestTypeJson,
		"urlencoded": OAuthRequestTypeUrlencoded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OAuthRequestType(input)
	return &out, nil
}

type TileTextSizeUnit string

const (
	TileTextSizeUnitPt TileTextSizeUnit = "pt"
)

func PossibleValuesForTileTextSizeUnit() []string {
	return []string{
		string(TileTextSizeUnitPt),
	}
}

func (s *TileTextSizeUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTileTextSizeUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTileTextSizeUnit(input string) (*TileTextSizeUnit, error) {
	vals := map[string]TileTextSizeUnit{
		"pt": TileTextSizeUnitPt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TileTextSizeUnit(input)
	return &out, nil
}

type TileTimeRangeDuration string

const (
	TileTimeRangeDurationPOneD        TileTimeRangeDuration = "P1D"
	TileTimeRangeDurationPOneW        TileTimeRangeDuration = "P1W"
	TileTimeRangeDurationPTOneH       TileTimeRangeDuration = "PT1H"
	TileTimeRangeDurationPTOneTwoH    TileTimeRangeDuration = "PT12H"
	TileTimeRangeDurationPTThreeZeroM TileTimeRangeDuration = "PT30M"
	TileTimeRangeDurationPThreeZeroD  TileTimeRangeDuration = "P30D"
)

func PossibleValuesForTileTimeRangeDuration() []string {
	return []string{
		string(TileTimeRangeDurationPOneD),
		string(TileTimeRangeDurationPOneW),
		string(TileTimeRangeDurationPTOneH),
		string(TileTimeRangeDurationPTOneTwoH),
		string(TileTimeRangeDurationPTThreeZeroM),
		string(TileTimeRangeDurationPThreeZeroD),
	}
}

func (s *TileTimeRangeDuration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTileTimeRangeDuration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTileTimeRangeDuration(input string) (*TileTimeRangeDuration, error) {
	vals := map[string]TileTimeRangeDuration{
		"p1d":   TileTimeRangeDurationPOneD,
		"p1w":   TileTimeRangeDurationPOneW,
		"pt1h":  TileTimeRangeDurationPTOneH,
		"pt12h": TileTimeRangeDurationPTOneTwoH,
		"pt30m": TileTimeRangeDurationPTThreeZeroM,
		"p30d":  TileTimeRangeDurationPThreeZeroD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TileTimeRangeDuration(input)
	return &out, nil
}

type TileTimeRangeResolution string

const (
	TileTimeRangeResolutionPOneD        TileTimeRangeResolution = "P1D"
	TileTimeRangeResolutionPOneW        TileTimeRangeResolution = "P1W"
	TileTimeRangeResolutionPTFiveM      TileTimeRangeResolution = "PT5M"
	TileTimeRangeResolutionPTOneH       TileTimeRangeResolution = "PT1H"
	TileTimeRangeResolutionPTOneM       TileTimeRangeResolution = "PT1M"
	TileTimeRangeResolutionPTOneTwoH    TileTimeRangeResolution = "PT12H"
	TileTimeRangeResolutionPTOneZeroM   TileTimeRangeResolution = "PT10M"
	TileTimeRangeResolutionPTThreeZeroM TileTimeRangeResolution = "PT30M"
)

func PossibleValuesForTileTimeRangeResolution() []string {
	return []string{
		string(TileTimeRangeResolutionPOneD),
		string(TileTimeRangeResolutionPOneW),
		string(TileTimeRangeResolutionPTFiveM),
		string(TileTimeRangeResolutionPTOneH),
		string(TileTimeRangeResolutionPTOneM),
		string(TileTimeRangeResolutionPTOneTwoH),
		string(TileTimeRangeResolutionPTOneZeroM),
		string(TileTimeRangeResolutionPTThreeZeroM),
	}
}

func (s *TileTimeRangeResolution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTileTimeRangeResolution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTileTimeRangeResolution(input string) (*TileTimeRangeResolution, error) {
	vals := map[string]TileTimeRangeResolution{
		"p1d":   TileTimeRangeResolutionPOneD,
		"p1w":   TileTimeRangeResolutionPOneW,
		"pt5m":  TileTimeRangeResolutionPTFiveM,
		"pt1h":  TileTimeRangeResolutionPTOneH,
		"pt1m":  TileTimeRangeResolutionPTOneM,
		"pt12h": TileTimeRangeResolutionPTOneTwoH,
		"pt10m": TileTimeRangeResolutionPTOneZeroM,
		"pt30m": TileTimeRangeResolutionPTThreeZeroM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TileTimeRangeResolution(input)
	return &out, nil
}
