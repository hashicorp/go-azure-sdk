package fileimports

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteStatus string

const (
	DeleteStatusDeleted     DeleteStatus = "Deleted"
	DeleteStatusNotDeleted  DeleteStatus = "NotDeleted"
	DeleteStatusUnspecified DeleteStatus = "Unspecified"
)

func PossibleValuesForDeleteStatus() []string {
	return []string{
		string(DeleteStatusDeleted),
		string(DeleteStatusNotDeleted),
		string(DeleteStatusUnspecified),
	}
}

func parseDeleteStatus(input string) (*DeleteStatus, error) {
	vals := map[string]DeleteStatus{
		"deleted":     DeleteStatusDeleted,
		"notdeleted":  DeleteStatusNotDeleted,
		"unspecified": DeleteStatusUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteStatus(input)
	return &out, nil
}

type FileFormat string

const (
	FileFormatCSV         FileFormat = "CSV"
	FileFormatJSON        FileFormat = "JSON"
	FileFormatUnspecified FileFormat = "Unspecified"
)

func PossibleValuesForFileFormat() []string {
	return []string{
		string(FileFormatCSV),
		string(FileFormatJSON),
		string(FileFormatUnspecified),
	}
}

func parseFileFormat(input string) (*FileFormat, error) {
	vals := map[string]FileFormat{
		"csv":         FileFormatCSV,
		"json":        FileFormatJSON,
		"unspecified": FileFormatUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileFormat(input)
	return &out, nil
}

type FileImportContentType string

const (
	FileImportContentTypeBasicIndicator FileImportContentType = "BasicIndicator"
	FileImportContentTypeStixIndicator  FileImportContentType = "StixIndicator"
	FileImportContentTypeUnspecified    FileImportContentType = "Unspecified"
)

func PossibleValuesForFileImportContentType() []string {
	return []string{
		string(FileImportContentTypeBasicIndicator),
		string(FileImportContentTypeStixIndicator),
		string(FileImportContentTypeUnspecified),
	}
}

func parseFileImportContentType(input string) (*FileImportContentType, error) {
	vals := map[string]FileImportContentType{
		"basicindicator": FileImportContentTypeBasicIndicator,
		"stixindicator":  FileImportContentTypeStixIndicator,
		"unspecified":    FileImportContentTypeUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileImportContentType(input)
	return &out, nil
}

type FileImportState string

const (
	FileImportStateFatalError         FileImportState = "FatalError"
	FileImportStateInProgress         FileImportState = "InProgress"
	FileImportStateIngested           FileImportState = "Ingested"
	FileImportStateIngestedWithErrors FileImportState = "IngestedWithErrors"
	FileImportStateInvalid            FileImportState = "Invalid"
	FileImportStateUnspecified        FileImportState = "Unspecified"
	FileImportStateWaitingForUpload   FileImportState = "WaitingForUpload"
)

func PossibleValuesForFileImportState() []string {
	return []string{
		string(FileImportStateFatalError),
		string(FileImportStateInProgress),
		string(FileImportStateIngested),
		string(FileImportStateIngestedWithErrors),
		string(FileImportStateInvalid),
		string(FileImportStateUnspecified),
		string(FileImportStateWaitingForUpload),
	}
}

func parseFileImportState(input string) (*FileImportState, error) {
	vals := map[string]FileImportState{
		"fatalerror":         FileImportStateFatalError,
		"inprogress":         FileImportStateInProgress,
		"ingested":           FileImportStateIngested,
		"ingestedwitherrors": FileImportStateIngestedWithErrors,
		"invalid":            FileImportStateInvalid,
		"unspecified":        FileImportStateUnspecified,
		"waitingforupload":   FileImportStateWaitingForUpload,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileImportState(input)
	return &out, nil
}

type IngestionMode string

const (
	IngestionModeIngestAnyValidRecords   IngestionMode = "IngestAnyValidRecords"
	IngestionModeIngestOnlyIfAllAreValid IngestionMode = "IngestOnlyIfAllAreValid"
	IngestionModeUnspecified             IngestionMode = "Unspecified"
)

func PossibleValuesForIngestionMode() []string {
	return []string{
		string(IngestionModeIngestAnyValidRecords),
		string(IngestionModeIngestOnlyIfAllAreValid),
		string(IngestionModeUnspecified),
	}
}

func parseIngestionMode(input string) (*IngestionMode, error) {
	vals := map[string]IngestionMode{
		"ingestanyvalidrecords":   IngestionModeIngestAnyValidRecords,
		"ingestonlyifallarevalid": IngestionModeIngestOnlyIfAllAreValid,
		"unspecified":             IngestionModeUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IngestionMode(input)
	return &out, nil
}
