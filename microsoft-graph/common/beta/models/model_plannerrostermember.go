package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRosterMember struct {
	Id        *string   `json:"id,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
	Roles     *[]string `json:"roles,omitempty"`
	TenantId  *string   `json:"tenantId,omitempty"`
	UserId    *string   `json:"userId,omitempty"`
}
