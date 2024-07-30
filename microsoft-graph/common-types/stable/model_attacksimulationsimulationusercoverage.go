package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationSimulationUserCoverage struct {
	AttackSimulationUser     *AttackSimulationUser `json:"attackSimulationUser,omitempty"`
	ClickCount               *int64                `json:"clickCount,omitempty"`
	CompromisedCount         *int64                `json:"compromisedCount,omitempty"`
	LatestSimulationDateTime *string               `json:"latestSimulationDateTime,omitempty"`
	ODataType                *string               `json:"@odata.type,omitempty"`
	SimulationCount          *int64                `json:"simulationCount,omitempty"`
}
