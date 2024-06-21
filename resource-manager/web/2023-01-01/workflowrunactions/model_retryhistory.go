package workflowrunactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetryHistory struct {
	ClientRequestId  *string        `json:"clientRequestId,omitempty"`
	Code             *string        `json:"code,omitempty"`
	EndTime          *string        `json:"endTime,omitempty"`
	Error            *ErrorResponse `json:"error,omitempty"`
	ServiceRequestId *string        `json:"serviceRequestId,omitempty"`
	StartTime        *string        `json:"startTime,omitempty"`
}
