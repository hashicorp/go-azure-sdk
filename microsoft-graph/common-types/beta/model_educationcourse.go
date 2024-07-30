package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationCourse struct {
	CourseNumber *string `json:"courseNumber,omitempty"`
	Description  *string `json:"description,omitempty"`
	DisplayName  *string `json:"displayName,omitempty"`
	ExternalId   *string `json:"externalId,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Subject      *string `json:"subject,omitempty"`
}
