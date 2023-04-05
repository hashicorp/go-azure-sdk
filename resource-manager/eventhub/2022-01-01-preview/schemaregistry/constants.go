package schemaregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaCompatibility string

const (
	SchemaCompatibilityBackward SchemaCompatibility = "Backward"
	SchemaCompatibilityForward  SchemaCompatibility = "Forward"
	SchemaCompatibilityNone     SchemaCompatibility = "None"
)

func PossibleValuesForSchemaCompatibility() []string {
	return []string{
		string(SchemaCompatibilityBackward),
		string(SchemaCompatibilityForward),
		string(SchemaCompatibilityNone),
	}
}

type SchemaType string

const (
	SchemaTypeAvro    SchemaType = "Avro"
	SchemaTypeUnknown SchemaType = "Unknown"
)

func PossibleValuesForSchemaType() []string {
	return []string{
		string(SchemaTypeAvro),
		string(SchemaTypeUnknown),
	}
}
