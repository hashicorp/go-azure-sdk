package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyGroupAssignment struct {
	DeviceCompliancePolicy *DeviceCompliancePolicy `json:"deviceCompliancePolicy,omitempty"`
	ExcludeGroup           *bool                   `json:"excludeGroup,omitempty"`
	Id                     *string                 `json:"id,omitempty"`
	ODataType              *string                 `json:"@odata.type,omitempty"`
	TargetGroupId          *string                 `json:"targetGroupId,omitempty"`
}
