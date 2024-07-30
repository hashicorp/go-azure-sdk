package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermission struct {
	ODataType       *string                                           `json:"@odata.type,omitempty"`
	PermissionType  *TeamsAppResourceSpecificPermissionPermissionType `json:"permissionType,omitempty"`
	PermissionValue *string                                           `json:"permissionValue,omitempty"`
}
