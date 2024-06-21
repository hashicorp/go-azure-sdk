package sessionhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostHealthCheckFailureDetails struct {
	ErrorCode               *int64  `json:"errorCode,omitempty"`
	LastHealthCheckDateTime *string `json:"lastHealthCheckDateTime,omitempty"`
	Message                 *string `json:"message,omitempty"`
}
