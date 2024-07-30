package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingEventInfo struct {
	DisplayName                 *string                                    `json:"displayName,omitempty"`
	LatestTrainingStatus        *UserTrainingEventInfoLatestTrainingStatus `json:"latestTrainingStatus,omitempty"`
	ODataType                   *string                                    `json:"@odata.type,omitempty"`
	TrainingAssignedProperties  *UserTrainingContentEventInfo              `json:"trainingAssignedProperties,omitempty"`
	TrainingCompletedProperties *UserTrainingContentEventInfo              `json:"trainingCompletedProperties,omitempty"`
	TrainingUpdatedProperties   *UserTrainingContentEventInfo              `json:"trainingUpdatedProperties,omitempty"`
}
