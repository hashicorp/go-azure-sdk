package targetcomputesizes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeSizeErrorDetails struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}
