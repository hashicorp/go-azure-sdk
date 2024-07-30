package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinition struct {
	AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem         `json:"authorizationSystemInfo,omitempty"`
	IdentityInfo            *PermissionsDefinitionAuthorizationSystemIdentity `json:"identityInfo,omitempty"`
	ODataType               *string                                           `json:"@odata.type,omitempty"`
}
