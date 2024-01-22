package billingroledefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleDefinitionProperties struct {
	Description *string                         `json:"description,omitempty"`
	Permissions *[]BillingPermissionsProperties `json:"permissions,omitempty"`
	RoleName    *string                         `json:"roleName,omitempty"`
}
