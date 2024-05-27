package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentCommonProperties struct {
	DataTypeScope []string            `json:"dataTypeScope"`
	PrincipalId   string              `json:"principalId"`
	PrincipalType string              `json:"principalType"`
	Role          DataProductUserRole `json:"role"`
	RoleId        string              `json:"roleId"`
	UserName      string              `json:"userName"`
}
