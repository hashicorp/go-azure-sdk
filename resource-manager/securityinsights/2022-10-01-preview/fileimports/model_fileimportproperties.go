package fileimports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileImportProperties struct {
	ContentType             FileImportContentType `json:"contentType"`
	CreatedTimeUTC          *string               `json:"createdTimeUTC,omitempty"`
	ErrorFile               *FileMetadata         `json:"errorFile,omitempty"`
	ErrorsPreview           *[]ValidationError    `json:"errorsPreview,omitempty"`
	FilesValidUntilTimeUTC  *string               `json:"filesValidUntilTimeUTC,omitempty"`
	ImportFile              FileMetadata          `json:"importFile"`
	ImportValidUntilTimeUTC *string               `json:"importValidUntilTimeUTC,omitempty"`
	IngestedRecordCount     *int64                `json:"ingestedRecordCount,omitempty"`
	IngestionMode           IngestionMode         `json:"ingestionMode"`
	Source                  string                `json:"source"`
	State                   *FileImportState      `json:"state,omitempty"`
	TotalRecordCount        *int64                `json:"totalRecordCount,omitempty"`
	ValidRecordCount        *int64                `json:"validRecordCount,omitempty"`
}
