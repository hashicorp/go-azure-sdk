package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStageStatus struct {
	AfterGates           *[]UpdateRunGateStatus `json:"afterGates,omitempty"`
	AfterStageWaitStatus *WaitStatus            `json:"afterStageWaitStatus,omitempty"`
	BeforeGates          *[]UpdateRunGateStatus `json:"beforeGates,omitempty"`
	Groups               *[]UpdateGroupStatus   `json:"groups,omitempty"`
	MaxConcurrency       *int64                 `json:"maxConcurrency,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Status               *UpdateStatus          `json:"status,omitempty"`
}
