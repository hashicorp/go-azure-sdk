package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingEventsContent struct {
	AssignedTrainingsInfos     *[]AssignedTrainingInfo `json:"assignedTrainingsInfos,omitempty"`
	ODataType                  *string                 `json:"@odata.type,omitempty"`
	TrainingsAssignedUserCount *int64                  `json:"trainingsAssignedUserCount,omitempty"`
}
