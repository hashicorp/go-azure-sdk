package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningSelfInitiatedCourse struct {
	CompletedDateTime        *string                            `json:"completedDateTime,omitempty"`
	CompletionPercentage     *int64                             `json:"completionPercentage,omitempty"`
	ExternalcourseActivityId *string                            `json:"externalcourseActivityId,omitempty"`
	Id                       *string                            `json:"id,omitempty"`
	LearnerUserId            *string                            `json:"learnerUserId,omitempty"`
	LearningContentId        *string                            `json:"learningContentId,omitempty"`
	LearningProviderId       *string                            `json:"learningProviderId,omitempty"`
	ODataType                *string                            `json:"@odata.type,omitempty"`
	StartedDateTime          *string                            `json:"startedDateTime,omitempty"`
	Status                   *LearningSelfInitiatedCourseStatus `json:"status,omitempty"`
}
