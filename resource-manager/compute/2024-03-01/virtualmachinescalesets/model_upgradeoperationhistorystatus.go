package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpgradeOperationHistoryStatus struct {
	Code      *UpgradeState `json:"code,omitempty"`
	EndTime   *string       `json:"endTime,omitempty"`
	StartTime *string       `json:"startTime,omitempty"`
}
