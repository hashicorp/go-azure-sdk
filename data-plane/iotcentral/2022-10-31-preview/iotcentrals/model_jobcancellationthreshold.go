package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobCancellationThreshold struct {
	Batch *bool                        `json:"batch,omitempty"`
	Type  JobCancellationThresholdType `json:"type"`
	Value float64                      `json:"value"`
}
