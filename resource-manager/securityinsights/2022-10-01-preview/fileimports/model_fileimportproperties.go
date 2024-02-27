package fileimports

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *FileImportProperties) GetCreatedTimeUTCAsTime() (*time.Time, error) {
	if o.CreatedTimeUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *FileImportProperties) SetCreatedTimeUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimeUTC = &formatted
}

func (o *FileImportProperties) GetFilesValidUntilTimeUTCAsTime() (*time.Time, error) {
	if o.FilesValidUntilTimeUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FilesValidUntilTimeUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *FileImportProperties) SetFilesValidUntilTimeUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FilesValidUntilTimeUTC = &formatted
}

func (o *FileImportProperties) GetImportValidUntilTimeUTCAsTime() (*time.Time, error) {
	if o.ImportValidUntilTimeUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ImportValidUntilTimeUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *FileImportProperties) SetImportValidUntilTimeUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ImportValidUntilTimeUTC = &formatted
}
