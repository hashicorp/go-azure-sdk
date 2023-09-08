package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfile struct {
	Assignments                    *[]AppleEnrollmentProfileAssignment                       `json:"assignments,omitempty"`
	AvailableEnrollmentTypeOptions *[]AppleOwnerTypeEnrollmentType                           `json:"availableEnrollmentTypeOptions,omitempty"`
	CreatedDateTime                *string                                                   `json:"createdDateTime,omitempty"`
	DefaultEnrollmentType          *AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType `json:"defaultEnrollmentType,omitempty"`
	Description                    *string                                                   `json:"description,omitempty"`
	DisplayName                    *string                                                   `json:"displayName,omitempty"`
	Id                             *string                                                   `json:"id,omitempty"`
	LastModifiedDateTime           *string                                                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                      *string                                                   `json:"@odata.type,omitempty"`
	Platform                       *AppleUserInitiatedEnrollmentProfilePlatform              `json:"platform,omitempty"`
	Priority                       *int64                                                    `json:"priority,omitempty"`
}
