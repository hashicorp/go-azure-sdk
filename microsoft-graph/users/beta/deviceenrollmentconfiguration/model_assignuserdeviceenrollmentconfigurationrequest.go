package deviceenrollmentconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignUserDeviceEnrollmentConfigurationRequest struct {
	EnrollmentConfigurationAssignments *[]EnrollmentConfigurationAssignment `json:"enrollmentConfigurationAssignments,omitempty"`
}
