package inputs

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
