package streamingjobs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMode string

const (
	AuthenticationModeConnectionString AuthenticationMode = "ConnectionString"
	AuthenticationModeMsi              AuthenticationMode = "Msi"
	AuthenticationModeUserToken        AuthenticationMode = "UserToken"
)

func PossibleValuesForAuthenticationMode() []string {
	return []string{
		string(AuthenticationModeConnectionString),
		string(AuthenticationModeMsi),
		string(AuthenticationModeUserToken),
	}
}

func parseAuthenticationMode(input string) (*AuthenticationMode, error) {
	vals := map[string]AuthenticationMode{
		"connectionstring": AuthenticationModeConnectionString,
		"msi":              AuthenticationModeMsi,
		"usertoken":        AuthenticationModeUserToken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMode(input)
	return &out, nil
}

type CompatibilityLevel string

const (
	CompatibilityLevelOnePointTwo  CompatibilityLevel = "1.2"
	CompatibilityLevelOnePointZero CompatibilityLevel = "1.0"
)

func PossibleValuesForCompatibilityLevel() []string {
	return []string{
		string(CompatibilityLevelOnePointTwo),
		string(CompatibilityLevelOnePointZero),
	}
}

func parseCompatibilityLevel(input string) (*CompatibilityLevel, error) {
	vals := map[string]CompatibilityLevel{
		"1.2": CompatibilityLevelOnePointTwo,
		"1.0": CompatibilityLevelOnePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompatibilityLevel(input)
	return &out, nil
}

type CompressionType string

const (
	CompressionTypeDeflate CompressionType = "Deflate"
	CompressionTypeGZip    CompressionType = "GZip"
	CompressionTypeNone    CompressionType = "None"
)

func PossibleValuesForCompressionType() []string {
	return []string{
		string(CompressionTypeDeflate),
		string(CompressionTypeGZip),
		string(CompressionTypeNone),
	}
}

func parseCompressionType(input string) (*CompressionType, error) {
	vals := map[string]CompressionType{
		"deflate": CompressionTypeDeflate,
		"gzip":    CompressionTypeGZip,
		"none":    CompressionTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompressionType(input)
	return &out, nil
}

type ContentStoragePolicy string

const (
	ContentStoragePolicyJobStorageAccount ContentStoragePolicy = "JobStorageAccount"
	ContentStoragePolicySystemAccount     ContentStoragePolicy = "SystemAccount"
)

func PossibleValuesForContentStoragePolicy() []string {
	return []string{
		string(ContentStoragePolicyJobStorageAccount),
		string(ContentStoragePolicySystemAccount),
	}
}

func parseContentStoragePolicy(input string) (*ContentStoragePolicy, error) {
	vals := map[string]ContentStoragePolicy{
		"jobstorageaccount": ContentStoragePolicyJobStorageAccount,
		"systemaccount":     ContentStoragePolicySystemAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentStoragePolicy(input)
	return &out, nil
}

type EventSerializationType string

const (
	EventSerializationTypeAvro    EventSerializationType = "Avro"
	EventSerializationTypeCsv     EventSerializationType = "Csv"
	EventSerializationTypeJson    EventSerializationType = "Json"
	EventSerializationTypeParquet EventSerializationType = "Parquet"
)

func PossibleValuesForEventSerializationType() []string {
	return []string{
		string(EventSerializationTypeAvro),
		string(EventSerializationTypeCsv),
		string(EventSerializationTypeJson),
		string(EventSerializationTypeParquet),
	}
}

func parseEventSerializationType(input string) (*EventSerializationType, error) {
	vals := map[string]EventSerializationType{
		"avro":    EventSerializationTypeAvro,
		"csv":     EventSerializationTypeCsv,
		"json":    EventSerializationTypeJson,
		"parquet": EventSerializationTypeParquet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSerializationType(input)
	return &out, nil
}

type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust EventsOutOfOrderPolicy = "Adjust"
	EventsOutOfOrderPolicyDrop   EventsOutOfOrderPolicy = "Drop"
)

func PossibleValuesForEventsOutOfOrderPolicy() []string {
	return []string{
		string(EventsOutOfOrderPolicyAdjust),
		string(EventsOutOfOrderPolicyDrop),
	}
}

func parseEventsOutOfOrderPolicy(input string) (*EventsOutOfOrderPolicy, error) {
	vals := map[string]EventsOutOfOrderPolicy{
		"adjust": EventsOutOfOrderPolicyAdjust,
		"drop":   EventsOutOfOrderPolicyDrop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventsOutOfOrderPolicy(input)
	return &out, nil
}

type JobType string

const (
	JobTypeCloud JobType = "Cloud"
	JobTypeEdge  JobType = "Edge"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeCloud),
		string(JobTypeEdge),
	}
}

func parseJobType(input string) (*JobType, error) {
	vals := map[string]JobType{
		"cloud": JobTypeCloud,
		"edge":  JobTypeEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
}

type OutputErrorPolicy string

const (
	OutputErrorPolicyDrop OutputErrorPolicy = "Drop"
	OutputErrorPolicyStop OutputErrorPolicy = "Stop"
)

func PossibleValuesForOutputErrorPolicy() []string {
	return []string{
		string(OutputErrorPolicyDrop),
		string(OutputErrorPolicyStop),
	}
}

func parseOutputErrorPolicy(input string) (*OutputErrorPolicy, error) {
	vals := map[string]OutputErrorPolicy{
		"drop": OutputErrorPolicyDrop,
		"stop": OutputErrorPolicyStop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputErrorPolicy(input)
	return &out, nil
}

type OutputStartMode string

const (
	OutputStartModeCustomTime          OutputStartMode = "CustomTime"
	OutputStartModeJobStartTime        OutputStartMode = "JobStartTime"
	OutputStartModeLastOutputEventTime OutputStartMode = "LastOutputEventTime"
)

func PossibleValuesForOutputStartMode() []string {
	return []string{
		string(OutputStartModeCustomTime),
		string(OutputStartModeJobStartTime),
		string(OutputStartModeLastOutputEventTime),
	}
}

func parseOutputStartMode(input string) (*OutputStartMode, error) {
	vals := map[string]OutputStartMode{
		"customtime":          OutputStartModeCustomTime,
		"jobstarttime":        OutputStartModeJobStartTime,
		"lastoutputeventtime": OutputStartModeLastOutputEventTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputStartMode(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameStandard),
	}
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"standard": SkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}
