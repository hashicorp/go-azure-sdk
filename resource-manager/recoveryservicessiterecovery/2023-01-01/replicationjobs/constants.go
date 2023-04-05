package replicationjobs

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
