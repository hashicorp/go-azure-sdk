package schema

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaType string

const (
	SchemaTypeJson SchemaType = "json"
	SchemaTypeXml  SchemaType = "xml"
)

func PossibleValuesForSchemaType() []string {
	return []string{
		string(SchemaTypeJson),
		string(SchemaTypeXml),
	}
}
