package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCManagementGroupAssignmentTarget struct {
	GroupId       *string `json:"groupId,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	ServicePlanId *string `json:"servicePlanId,omitempty"`
}
