package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsIdentitySource struct {
	AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem `json:"authorizationSystemInfo,omitempty"`
	ODataType               *string                                   `json:"@odata.type,omitempty"`
}
