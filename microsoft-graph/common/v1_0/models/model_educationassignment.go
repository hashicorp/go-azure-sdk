package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignment struct {
	AddToCalendarAction                     *EducationAssignmentAddToCalendarAction `json:"addToCalendarAction,omitempty"`
	AddedStudentAction                      *EducationAssignmentAddedStudentAction  `json:"addedStudentAction,omitempty"`
	AllowLateSubmissions                    *bool                                   `json:"allowLateSubmissions,omitempty"`
	AllowStudentsToAddResourcesToSubmission *bool                                   `json:"allowStudentsToAddResourcesToSubmission,omitempty"`
	AssignDateTime                          *string                                 `json:"assignDateTime,omitempty"`
	AssignTo                                *EducationAssignmentRecipient           `json:"assignTo,omitempty"`
	AssignedDateTime                        *string                                 `json:"assignedDateTime,omitempty"`
	Categories                              *[]EducationCategory                    `json:"categories,omitempty"`
	ClassId                                 *string                                 `json:"classId,omitempty"`
	CloseDateTime                           *string                                 `json:"closeDateTime,omitempty"`
	CreatedBy                               *IdentitySet                            `json:"createdBy,omitempty"`
	CreatedDateTime                         *string                                 `json:"createdDateTime,omitempty"`
	DisplayName                             *string                                 `json:"displayName,omitempty"`
	DueDateTime                             *string                                 `json:"dueDateTime,omitempty"`
	FeedbackResourcesFolderUrl              *string                                 `json:"feedbackResourcesFolderUrl,omitempty"`
	Grading                                 *EducationAssignmentGradeType           `json:"grading,omitempty"`
	Id                                      *string                                 `json:"id,omitempty"`
	Instructions                            *EducationItemBody                      `json:"instructions,omitempty"`
	LastModifiedBy                          *IdentitySet                            `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime                    *string                                 `json:"lastModifiedDateTime,omitempty"`
	NotificationChannelUrl                  *string                                 `json:"notificationChannelUrl,omitempty"`
	ODataType                               *string                                 `json:"@odata.type,omitempty"`
	Resources                               *[]EducationAssignmentResource          `json:"resources,omitempty"`
	ResourcesFolderUrl                      *string                                 `json:"resourcesFolderUrl,omitempty"`
	Rubric                                  *EducationRubric                        `json:"rubric,omitempty"`
	Status                                  *EducationAssignmentStatus              `json:"status,omitempty"`
	Submissions                             *[]EducationSubmission                  `json:"submissions,omitempty"`
	WebUrl                                  *string                                 `json:"webUrl,omitempty"`
}
