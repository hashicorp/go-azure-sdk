package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProgress struct {
	Completed *int64 `json:"completed,omitempty"`
	Failed    *int64 `json:"failed,omitempty"`
	Pending   *int64 `json:"pending,omitempty"`
	Total     *int64 `json:"total,omitempty"`
}
