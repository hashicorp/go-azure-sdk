package bot

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityResponseBody struct {
	AbsCode *string `json:"absCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Valid   *bool   `json:"valid,omitempty"`
}
