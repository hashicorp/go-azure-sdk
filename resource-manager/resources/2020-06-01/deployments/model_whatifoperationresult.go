package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WhatIfOperationResult struct {
	Error      *ErrorResponse             `json:"error"`
	Properties *WhatIfOperationProperties `json:"properties"`
	Status     *string                    `json:"status,omitempty"`
}
