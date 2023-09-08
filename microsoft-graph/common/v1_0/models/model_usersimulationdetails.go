package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSimulationDetails struct {
	AssignedTrainingsCount   *int64                     `json:"assignedTrainingsCount,omitempty"`
	CompletedTrainingsCount  *int64                     `json:"completedTrainingsCount,omitempty"`
	CompromisedDateTime      *string                    `json:"compromisedDateTime,omitempty"`
	InProgressTrainingsCount *int64                     `json:"inProgressTrainingsCount,omitempty"`
	IsCompromised            *bool                      `json:"isCompromised,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
	ReportedPhishDateTime    *string                    `json:"reportedPhishDateTime,omitempty"`
	SimulationEvents         *[]UserSimulationEventInfo `json:"simulationEvents,omitempty"`
	SimulationUser           *AttackSimulationUser      `json:"simulationUser,omitempty"`
	TrainingEvents           *[]UserTrainingEventInfo   `json:"trainingEvents,omitempty"`
}
