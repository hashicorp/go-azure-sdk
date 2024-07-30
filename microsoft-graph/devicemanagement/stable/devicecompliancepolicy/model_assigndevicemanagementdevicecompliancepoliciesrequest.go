package devicecompliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementDeviceCompliancePoliciesRequest struct {
	Assignments *[]DeviceCompliancePolicyAssignment `json:"assignments,omitempty"`
}
