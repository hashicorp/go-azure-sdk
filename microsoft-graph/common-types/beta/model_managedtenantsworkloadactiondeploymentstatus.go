package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionDeploymentStatus struct {
	ActionId               *string                                             `json:"actionId,omitempty"`
	DeployedPolicyId       *string                                             `json:"deployedPolicyId,omitempty"`
	Error                  *GenericError                                       `json:"error,omitempty"`
	ExcludeGroups          *[]string                                           `json:"excludeGroups,omitempty"`
	IncludeAllUsers        *bool                                               `json:"includeAllUsers,omitempty"`
	IncludeGroups          *[]string                                           `json:"includeGroups,omitempty"`
	LastDeploymentDateTime *string                                             `json:"lastDeploymentDateTime,omitempty"`
	ODataType              *string                                             `json:"@odata.type,omitempty"`
	Status                 *ManagedTenantsWorkloadActionDeploymentStatusStatus `json:"status,omitempty"`
}
