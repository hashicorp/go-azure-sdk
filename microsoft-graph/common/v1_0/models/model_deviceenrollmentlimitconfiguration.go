package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentLimitConfiguration struct {
	Assignments          *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`
	CreatedDateTime      *string                              `json:"createdDateTime,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	Limit                *int64                               `json:"limit,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	Priority             *int64                               `json:"priority,omitempty"`
	Version              *int64                               `json:"version,omitempty"`
}
