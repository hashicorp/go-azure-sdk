package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationReport struct {
	ODataType       *string                   `json:"@odata.type,omitempty"`
	Overview        *SimulationReportOverview `json:"overview,omitempty"`
	SimulationUsers *[]UserSimulationDetails  `json:"simulationUsers,omitempty"`
}
