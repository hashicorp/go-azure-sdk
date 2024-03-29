package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryTestingResult struct {
	Error     *ErrorError               `json:"error,omitempty"`
	OutputUri *string                   `json:"outputUri,omitempty"`
	Status    *QueryTestingResultStatus `json:"status,omitempty"`
}
