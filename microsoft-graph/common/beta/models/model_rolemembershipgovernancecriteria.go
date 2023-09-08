package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleMembershipGovernanceCriteria struct {
	ODataType      *string `json:"@odata.type,omitempty"`
	RoleId         *string `json:"roleId,omitempty"`
	RoleTemplateId *string `json:"roleTemplateId,omitempty"`
}
