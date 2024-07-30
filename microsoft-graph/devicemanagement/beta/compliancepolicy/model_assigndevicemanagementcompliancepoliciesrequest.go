package compliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementCompliancePoliciesRequest struct {
	Assignments *[]DeviceManagementConfigurationPolicyAssignment `json:"assignments,omitempty"`
}
