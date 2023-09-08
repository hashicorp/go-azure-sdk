package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentGrade struct {
	GradedBy       *IdentitySet `json:"gradedBy,omitempty"`
	GradedDateTime *string      `json:"gradedDateTime,omitempty"`
	ODataType      *string      `json:"@odata.type,omitempty"`
}
