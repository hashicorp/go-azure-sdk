package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningAssignment struct {
	AssignedDateTime         *string                           `json:"assignedDateTime,omitempty"`
	AssignerUserId           *string                           `json:"assignerUserId,omitempty"`
	AssignmentType           *LearningAssignmentAssignmentType `json:"assignmentType,omitempty"`
	CompletedDateTime        *string                           `json:"completedDateTime,omitempty"`
	CompletionPercentage     *int64                            `json:"completionPercentage,omitempty"`
	DueDateTime              *DateTimeTimeZone                 `json:"dueDateTime,omitempty"`
	ExternalcourseActivityId *string                           `json:"externalcourseActivityId,omitempty"`
	Id                       *string                           `json:"id,omitempty"`
	LearnerUserId            *string                           `json:"learnerUserId,omitempty"`
	LearningContentId        *string                           `json:"learningContentId,omitempty"`
	LearningProviderId       *string                           `json:"learningProviderId,omitempty"`
	Notes                    *ItemBody                         `json:"notes,omitempty"`
	ODataType                *string                           `json:"@odata.type,omitempty"`
	Status                   *LearningAssignmentStatus         `json:"status,omitempty"`
}
