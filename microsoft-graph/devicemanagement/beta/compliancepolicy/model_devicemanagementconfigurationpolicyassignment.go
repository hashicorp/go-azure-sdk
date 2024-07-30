package compliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyAssignment struct {
	Id        *string                                              `json:"id,omitempty"`
	ODataType *string                                              `json:"@odata.type,omitempty"`
	Source    *DeviceManagementConfigurationPolicyAssignmentSource `json:"source,omitempty"`
	SourceId  *string                                              `json:"sourceId,omitempty"`
	Target    *DeviceAndAppManagementAssignmentTarget              `json:"target,omitempty"`
}
