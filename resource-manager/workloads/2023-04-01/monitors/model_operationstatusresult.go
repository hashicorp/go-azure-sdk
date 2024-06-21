package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusResult struct {
	EndTime         *string                  `json:"endTime,omitempty"`
	Error           *ErrorDetail             `json:"error,omitempty"`
	Id              *string                  `json:"id,omitempty"`
	Name            *string                  `json:"name,omitempty"`
	Operations      *[]OperationStatusResult `json:"operations,omitempty"`
	PercentComplete *float64                 `json:"percentComplete,omitempty"`
	StartTime       *string                  `json:"startTime,omitempty"`
	Status          string                   `json:"status"`
}
