package hcpoperatoridentityrolesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperatorIdentityRoles struct {
	Name            string                   `json:"name"`
	Required        OperatorIdentityRequired `json:"required"`
	RoleDefinitions []RoleDefinition         `json:"roleDefinitions"`
}
