package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleOwnerTypeEnrollmentType struct {
	EnrollmentType *AppleOwnerTypeEnrollmentTypeEnrollmentType `json:"enrollmentType,omitempty"`
	ODataType      *string                                     `json:"@odata.type,omitempty"`
	OwnerType      *AppleOwnerTypeEnrollmentTypeOwnerType      `json:"ownerType,omitempty"`
}
