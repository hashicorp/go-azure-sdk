package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoMfaOnRoleActivationAlertIncident struct {
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	RoleDisplayName *string `json:"roleDisplayName,omitempty"`
	RoleTemplateId  *string `json:"roleTemplateId,omitempty"`
}
