package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobErrorDetails struct {
	Code         *string         `json:"code,omitempty"`
	ErrorDetails *[]JobErrorItem `json:"errorDetails,omitempty"`
	Message      *string         `json:"message,omitempty"`
}
