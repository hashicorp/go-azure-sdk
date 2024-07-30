package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperience struct {
	Goals                    *Goals                    `json:"goals,omitempty"`
	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`
	LearningProviders        *[]LearningProvider       `json:"learningProviders,omitempty"`
	ODataType                *string                   `json:"@odata.type,omitempty"`
}
