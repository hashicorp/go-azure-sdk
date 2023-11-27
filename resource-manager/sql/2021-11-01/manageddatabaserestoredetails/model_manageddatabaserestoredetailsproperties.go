package manageddatabaserestoredetails

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseRestoreDetailsProperties struct {
	BlockReason              *string   `json:"blockReason,omitempty"`
	CurrentRestoringFileName *string   `json:"currentRestoringFileName,omitempty"`
	LastRestoredFileName     *string   `json:"lastRestoredFileName,omitempty"`
	LastRestoredFileTime     *string   `json:"lastRestoredFileTime,omitempty"`
	LastUploadedFileName     *string   `json:"lastUploadedFileName,omitempty"`
	LastUploadedFileTime     *string   `json:"lastUploadedFileTime,omitempty"`
	NumberOfFilesDetected    *int64    `json:"numberOfFilesDetected,omitempty"`
	PercentCompleted         *float64  `json:"percentCompleted,omitempty"`
	Status                   *string   `json:"status,omitempty"`
	UnrestorableFiles        *[]string `json:"unrestorableFiles,omitempty"`
}

func (o *ManagedDatabaseRestoreDetailsProperties) GetLastRestoredFileTimeAsTime() (*time.Time, error) {
	if o.LastRestoredFileTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRestoredFileTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedDatabaseRestoreDetailsProperties) SetLastRestoredFileTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRestoredFileTime = &formatted
}

func (o *ManagedDatabaseRestoreDetailsProperties) GetLastUploadedFileTimeAsTime() (*time.Time, error) {
	if o.LastUploadedFileTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUploadedFileTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedDatabaseRestoreDetailsProperties) SetLastUploadedFileTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUploadedFileTime = &formatted
}
