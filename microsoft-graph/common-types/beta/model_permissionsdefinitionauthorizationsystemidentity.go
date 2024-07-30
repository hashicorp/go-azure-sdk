package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAuthorizationSystemIdentity struct {
	ExternalId   *string                                                       `json:"externalId,omitempty"`
	IdentityType *PermissionsDefinitionAuthorizationSystemIdentityIdentityType `json:"identityType,omitempty"`
	ODataType    *string                                                       `json:"@odata.type,omitempty"`
	Source       *PermissionsDefinitionIdentitySource                          `json:"source,omitempty"`
}
