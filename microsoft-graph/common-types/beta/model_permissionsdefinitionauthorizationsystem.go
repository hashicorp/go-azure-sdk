package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAuthorizationSystem struct {
	AuthorizationSystemId   *string `json:"authorizationSystemId,omitempty"`
	AuthorizationSystemType *string `json:"authorizationSystemType,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
}
