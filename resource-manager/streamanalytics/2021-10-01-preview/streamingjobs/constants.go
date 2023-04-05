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

type BlobWriteMode string

const (
	BlobWriteModeAppend BlobWriteMode = "Append"
	BlobWriteModeOnce   BlobWriteMode = "Once"
)

func PossibleValuesForBlobWriteMode() []string {
	return []string{
		string(BlobWriteModeAppend),
		string(BlobWriteModeOnce),
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

type EventGridEventSchemaType string

const (
	EventGridEventSchemaTypeCloudEventSchema     EventGridEventSchemaType = "CloudEventSchema"
	EventGridEventSchemaTypeEventGridEventSchema EventGridEventSchemaType = "EventGridEventSchema"
)

func PossibleValuesForEventGridEventSchemaType() []string {
	return []string{
		string(EventGridEventSchemaTypeCloudEventSchema),
		string(EventGridEventSchemaTypeEventGridEventSchema),
	}
}

type EventSerializationType string

const (
	EventSerializationTypeAvro      EventSerializationType = "Avro"
	EventSerializationTypeCsv       EventSerializationType = "Csv"
	EventSerializationTypeCustomClr EventSerializationType = "CustomClr"
	EventSerializationTypeDelta     EventSerializationType = "Delta"
	EventSerializationTypeJson      EventSerializationType = "Json"
	EventSerializationTypeParquet   EventSerializationType = "Parquet"
)

func PossibleValuesForEventSerializationType() []string {
	return []string{
		string(EventSerializationTypeAvro),
		string(EventSerializationTypeCsv),
		string(EventSerializationTypeCustomClr),
		string(EventSerializationTypeDelta),
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

type InputWatermarkMode string

const (
	InputWatermarkModeNone          InputWatermarkMode = "None"
	InputWatermarkModeReadWatermark InputWatermarkMode = "ReadWatermark"
)

func PossibleValuesForInputWatermarkMode() []string {
	return []string{
		string(InputWatermarkModeNone),
		string(InputWatermarkModeReadWatermark),
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

type OutputWatermarkMode string

const (
	OutputWatermarkModeNone                                OutputWatermarkMode = "None"
	OutputWatermarkModeSendCurrentPartitionWatermark       OutputWatermarkMode = "SendCurrentPartitionWatermark"
	OutputWatermarkModeSendLowestWatermarkAcrossPartitions OutputWatermarkMode = "SendLowestWatermarkAcrossPartitions"
)

func PossibleValuesForOutputWatermarkMode() []string {
	return []string{
		string(OutputWatermarkModeNone),
		string(OutputWatermarkModeSendCurrentPartitionWatermark),
		string(OutputWatermarkModeSendLowestWatermarkAcrossPartitions),
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

type ResourceType string

const (
	ResourceTypeMicrosoftPointStreamAnalyticsStreamingjobs ResourceType = "Microsoft.StreamAnalytics/streamingjobs"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeMicrosoftPointStreamAnalyticsStreamingjobs),
	}
}

type SkuCapacityScaleType string

const (
	SkuCapacityScaleTypeAutomatic SkuCapacityScaleType = "automatic"
	SkuCapacityScaleTypeManual    SkuCapacityScaleType = "manual"
	SkuCapacityScaleTypeNone      SkuCapacityScaleType = "none"
)

func PossibleValuesForSkuCapacityScaleType() []string {
	return []string{
		string(SkuCapacityScaleTypeAutomatic),
		string(SkuCapacityScaleTypeManual),
		string(SkuCapacityScaleTypeNone),
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

type UpdatableUdfRefreshType string

const (
	UpdatableUdfRefreshTypeBlocking    UpdatableUdfRefreshType = "Blocking"
	UpdatableUdfRefreshTypeNonblocking UpdatableUdfRefreshType = "Nonblocking"
)

func PossibleValuesForUpdatableUdfRefreshType() []string {
	return []string{
		string(UpdatableUdfRefreshTypeBlocking),
		string(UpdatableUdfRefreshTypeNonblocking),
	}
}

type UpdateMode string

const (
	UpdateModeRefreshable UpdateMode = "Refreshable"
	UpdateModeStatic      UpdateMode = "Static"
)

func PossibleValuesForUpdateMode() []string {
	return []string{
		string(UpdateModeRefreshable),
		string(UpdateModeStatic),
	}
}
