package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingStatusInfo struct {
	AssignedDateTime   *string                               `json:"assignedDateTime,omitempty"`
	CompletionDateTime *string                               `json:"completionDateTime,omitempty"`
	DisplayName        *string                               `json:"displayName,omitempty"`
	ODataType          *string                               `json:"@odata.type,omitempty"`
	TrainingStatus     *UserTrainingStatusInfoTrainingStatus `json:"trainingStatus,omitempty"`
}
