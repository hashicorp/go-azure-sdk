package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleAssignment struct {
	ExpirationDateTime *string         `json:"expirationDateTime,omitempty"`
	Id                 *string         `json:"id,omitempty"`
	IsElevated         *bool           `json:"isElevated,omitempty"`
	ODataType          *string         `json:"@odata.type,omitempty"`
	ResultMessage      *string         `json:"resultMessage,omitempty"`
	RoleId             *string         `json:"roleId,omitempty"`
	RoleInfo           *PrivilegedRole `json:"roleInfo,omitempty"`
	UserId             *string         `json:"userId,omitempty"`
}
