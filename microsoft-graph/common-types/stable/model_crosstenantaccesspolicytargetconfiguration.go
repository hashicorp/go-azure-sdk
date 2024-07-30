package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetConfiguration struct {
	AccessType *CrossTenantAccessPolicyTargetConfigurationAccessType `json:"accessType,omitempty"`
	ODataType  *string                                               `json:"@odata.type,omitempty"`
	Targets    *[]CrossTenantAccessPolicyTarget                      `json:"targets,omitempty"`
}
