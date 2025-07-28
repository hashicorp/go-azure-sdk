package dtls

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationResult struct {
	Error      *OperationError `json:"error,omitempty"`
	Status     *string         `json:"status,omitempty"`
	StatusCode *HTTPStatusCode `json:"statusCode,omitempty"`
}
