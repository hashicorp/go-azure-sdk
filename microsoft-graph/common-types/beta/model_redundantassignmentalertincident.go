package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedundantAssignmentAlertIncident struct {
	AssigneeDisplayName       *string `json:"assigneeDisplayName,omitempty"`
	AssigneeId                *string `json:"assigneeId,omitempty"`
	AssigneeUserPrincipalName *string `json:"assigneeUserPrincipalName,omitempty"`
	Id                        *string `json:"id,omitempty"`
	LastActivationDateTime    *string `json:"lastActivationDateTime,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	RoleDefinitionId          *string `json:"roleDefinitionId,omitempty"`
	RoleDisplayName           *string `json:"roleDisplayName,omitempty"`
	RoleTemplateId            *string `json:"roleTemplateId,omitempty"`
}
