package replicationjobs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportJobOutputSerializationType string

const (
	ExportJobOutputSerializationTypeExcel ExportJobOutputSerializationType = "Excel"
	ExportJobOutputSerializationTypeJson  ExportJobOutputSerializationType = "Json"
	ExportJobOutputSerializationTypeXml   ExportJobOutputSerializationType = "Xml"
)

func PossibleValuesForExportJobOutputSerializationType() []string {
	return []string{
		string(ExportJobOutputSerializationTypeExcel),
		string(ExportJobOutputSerializationTypeJson),
		string(ExportJobOutputSerializationTypeXml),
	}
}

func parseExportJobOutputSerializationType(input string) (*ExportJobOutputSerializationType, error) {
	vals := map[string]ExportJobOutputSerializationType{
		"excel": ExportJobOutputSerializationTypeExcel,
		"json":  ExportJobOutputSerializationTypeJson,
		"xml":   ExportJobOutputSerializationTypeXml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportJobOutputSerializationType(input)
	return &out, nil
}
