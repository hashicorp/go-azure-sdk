package fleetupdatestrategies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStage struct {
	AfterGates              *[]GateConfiguration `json:"afterGates,omitempty"`
	AfterStageWaitInSeconds *int64               `json:"afterStageWaitInSeconds,omitempty"`
	BeforeGates             *[]GateConfiguration `json:"beforeGates,omitempty"`
	Groups                  *[]UpdateGroup       `json:"groups,omitempty"`
	MaxConcurrency          *string              `json:"maxConcurrency,omitempty"`
	Name                    string               `json:"name"`
}
