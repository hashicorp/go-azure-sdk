package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRoleBasedRule struct {
	DefaultRule  *string                           `json:"defaultRule,omitempty"`
	ODataType    *string                           `json:"@odata.type,omitempty"`
	PropertyRule *PlannerTaskPropertyRule          `json:"propertyRule,omitempty"`
	Role         *PlannerTaskConfigurationRoleBase `json:"role,omitempty"`
}
