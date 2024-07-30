package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaults struct {
	AddToCalendarAction    *EducationAssignmentDefaultsAddToCalendarAction `json:"addToCalendarAction,omitempty"`
	AddedStudentAction     *EducationAssignmentDefaultsAddedStudentAction  `json:"addedStudentAction,omitempty"`
	DueTime                *string                                         `json:"dueTime,omitempty"`
	Id                     *string                                         `json:"id,omitempty"`
	NotificationChannelUrl *string                                         `json:"notificationChannelUrl,omitempty"`
	ODataType              *string                                         `json:"@odata.type,omitempty"`
}
