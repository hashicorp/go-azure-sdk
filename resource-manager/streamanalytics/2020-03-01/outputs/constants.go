package outputs

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
