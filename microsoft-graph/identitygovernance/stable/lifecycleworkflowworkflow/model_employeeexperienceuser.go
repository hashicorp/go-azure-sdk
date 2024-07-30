package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceUser struct {
	Id                       *string                   `json:"id,omitempty"`
	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`
	ODataType                *string                   `json:"@odata.type,omitempty"`
}
