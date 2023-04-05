package streamingjobs

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

type Encoding string

const (
	EncodingUTFEight Encoding = "UTF8"
)

func PossibleValuesForEncoding() []string {
	return []string{
		string(EncodingUTFEight),
	}
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

type JsonOutputSerializationFormat string

const (
	JsonOutputSerializationFormatArray         JsonOutputSerializationFormat = "Array"
	JsonOutputSerializationFormatLineSeparated JsonOutputSerializationFormat = "LineSeparated"
)

func PossibleValuesForJsonOutputSerializationFormat() []string {
	return []string{
		string(JsonOutputSerializationFormatArray),
		string(JsonOutputSerializationFormatLineSeparated),
	}
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

type RefreshType string

const (
	RefreshTypeRefreshPeriodicallyWithDelta RefreshType = "RefreshPeriodicallyWithDelta"
	RefreshTypeRefreshPeriodicallyWithFull  RefreshType = "RefreshPeriodicallyWithFull"
	RefreshTypeStatic                       RefreshType = "Static"
)

func PossibleValuesForRefreshType() []string {
	return []string{
		string(RefreshTypeRefreshPeriodicallyWithDelta),
		string(RefreshTypeRefreshPeriodicallyWithFull),
		string(RefreshTypeStatic),
	}
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
