package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationDetails struct {
	DataSetId    *string      `json:"dataSetId,omitempty"`
	DataSetType  *DataSetType `json:"dataSetType,omitempty"`
	DurationMs   *int64       `json:"durationMs,omitempty"`
	EndTime      *string      `json:"endTime,omitempty"`
	FilesRead    *int64       `json:"filesRead,omitempty"`
	FilesWritten *int64       `json:"filesWritten,omitempty"`
	Message      *string      `json:"message,omitempty"`
	Name         *string      `json:"name,omitempty"`
	RowsCopied   *int64       `json:"rowsCopied,omitempty"`
	RowsRead     *int64       `json:"rowsRead,omitempty"`
	SizeRead     *int64       `json:"sizeRead,omitempty"`
	SizeWritten  *int64       `json:"sizeWritten,omitempty"`
	StartTime    *string      `json:"startTime,omitempty"`
	Status       *string      `json:"status,omitempty"`
	VCore        *int64       `json:"vCore,omitempty"`
}
