package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationCustomizations struct {
	ODataType         *string                                `json:"@odata.type,omitempty"`
	School            *EducationSynchronizationCustomization `json:"school,omitempty"`
	Section           *EducationSynchronizationCustomization `json:"section,omitempty"`
	Student           *EducationSynchronizationCustomization `json:"student,omitempty"`
	StudentEnrollment *EducationSynchronizationCustomization `json:"studentEnrollment,omitempty"`
	Teacher           *EducationSynchronizationCustomization `json:"teacher,omitempty"`
	TeacherRoster     *EducationSynchronizationCustomization `json:"teacherRoster,omitempty"`
}
