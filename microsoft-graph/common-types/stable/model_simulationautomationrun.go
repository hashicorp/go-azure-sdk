package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomationRun struct {
	EndDateTime   *string                        `json:"endDateTime,omitempty"`
	Id            *string                        `json:"id,omitempty"`
	ODataType     *string                        `json:"@odata.type,omitempty"`
	SimulationId  *string                        `json:"simulationId,omitempty"`
	StartDateTime *string                        `json:"startDateTime,omitempty"`
	Status        *SimulationAutomationRunStatus `json:"status,omitempty"`
}
