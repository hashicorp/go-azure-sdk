package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationRoot struct {
	Id                    *string                 `json:"id,omitempty"`
	ODataType             *string                 `json:"@odata.type,omitempty"`
	SimulationAutomations *[]SimulationAutomation `json:"simulationAutomations,omitempty"`
	Simulations           *[]Simulation           `json:"simulations,omitempty"`
}
