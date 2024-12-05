package inferencegroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActualCapacityInfo struct {
	Failed            *int64 `json:"failed,omitempty"`
	OutdatedFailed    *int64 `json:"outdatedFailed,omitempty"`
	OutdatedSucceeded *int64 `json:"outdatedSucceeded,omitempty"`
	Succeeded         *int64 `json:"succeeded,omitempty"`
	Total             *int64 `json:"total,omitempty"`
}
