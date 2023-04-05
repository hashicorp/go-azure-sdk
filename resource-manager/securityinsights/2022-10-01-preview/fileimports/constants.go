package fileimports

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
