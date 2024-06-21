package virtualmachinescalesetrollingupgrades

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RollingUpgradeRunningStatus struct {
	Code           *RollingUpgradeStatusCode `json:"code,omitempty"`
	LastAction     *RollingUpgradeActionType `json:"lastAction,omitempty"`
	LastActionTime *string                   `json:"lastActionTime,omitempty"`
	StartTime      *string                   `json:"startTime,omitempty"`
}
