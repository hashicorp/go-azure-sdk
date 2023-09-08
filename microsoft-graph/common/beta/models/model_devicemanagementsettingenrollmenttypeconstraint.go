package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingEnrollmentTypeConstraint struct {
	EnrollmentTypes *[]string `json:"enrollmentTypes,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
}
