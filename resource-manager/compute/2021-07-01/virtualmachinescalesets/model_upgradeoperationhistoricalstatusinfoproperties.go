package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpgradeOperationHistoricalStatusInfoProperties struct {
	Error                *ApiError                      `json:"error"`
	Progress             *RollingUpgradeProgressInfo    `json:"progress"`
	RollbackInfo         *RollbackStatusInfo            `json:"rollbackInfo"`
	RunningStatus        *UpgradeOperationHistoryStatus `json:"runningStatus"`
	StartedBy            *UpgradeOperationInvoker       `json:"startedBy,omitempty"`
	TargetImageReference *ImageReference                `json:"targetImageReference"`
}
