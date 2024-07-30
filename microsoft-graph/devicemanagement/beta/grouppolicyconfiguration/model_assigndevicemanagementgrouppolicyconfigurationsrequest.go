package grouppolicyconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementGroupPolicyConfigurationsRequest struct {
	Assignments *[]GroupPolicyConfigurationAssignment `json:"assignments,omitempty"`
}
