package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MECRoleProperties struct {
	ConnectionString   *AsymmetricEncryptedSecret `json:"connectionString,omitempty"`
	ControllerEndpoint *string                    `json:"controllerEndpoint,omitempty"`
	ResourceUniqueId   *string                    `json:"resourceUniqueId,omitempty"`
	RoleStatus         RoleStatus                 `json:"roleStatus"`
}
