package databasemigrationssqlvm

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyProgressDetails struct {
	CopyDuration       *int64   `json:"copyDuration,omitempty"`
	CopyStart          *string  `json:"copyStart,omitempty"`
	CopyThroughput     *float64 `json:"copyThroughput,omitempty"`
	DataRead           *int64   `json:"dataRead,omitempty"`
	DataWritten        *int64   `json:"dataWritten,omitempty"`
	ParallelCopyType   *string  `json:"parallelCopyType,omitempty"`
	RowsCopied         *int64   `json:"rowsCopied,omitempty"`
	RowsRead           *int64   `json:"rowsRead,omitempty"`
	Status             *string  `json:"status,omitempty"`
	TableName          *string  `json:"tableName,omitempty"`
	UsedParallelCopies *int64   `json:"usedParallelCopies,omitempty"`
}

func (o *CopyProgressDetails) GetCopyStartAsTime() (*time.Time, error) {
	if o.CopyStart == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CopyStart, "2006-01-02T15:04:05Z07:00")
}

func (o *CopyProgressDetails) SetCopyStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CopyStart = &formatted
}
