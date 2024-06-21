package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportRunProperties struct {
	Error               *ErrorDetails           `json:"error,omitempty"`
	ExecutionType       *ExecutionType          `json:"executionType,omitempty"`
	FileName            *string                 `json:"fileName,omitempty"`
	ProcessingEndTime   *string                 `json:"processingEndTime,omitempty"`
	ProcessingStartTime *string                 `json:"processingStartTime,omitempty"`
	RunSettings         *CommonExportProperties `json:"runSettings,omitempty"`
	Status              *ExecutionStatus        `json:"status,omitempty"`
	SubmittedBy         *string                 `json:"submittedBy,omitempty"`
	SubmittedTime       *string                 `json:"submittedTime,omitempty"`
}
