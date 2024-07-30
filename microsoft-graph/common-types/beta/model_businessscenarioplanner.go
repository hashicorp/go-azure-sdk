package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioPlanner struct {
	Id                *string                   `json:"id,omitempty"`
	ODataType         *string                   `json:"@odata.type,omitempty"`
	PlanConfiguration *PlannerPlanConfiguration `json:"planConfiguration,omitempty"`
	TaskConfiguration *PlannerTaskConfiguration `json:"taskConfiguration,omitempty"`
	Tasks             *[]BusinessScenarioTask   `json:"tasks,omitempty"`
}
