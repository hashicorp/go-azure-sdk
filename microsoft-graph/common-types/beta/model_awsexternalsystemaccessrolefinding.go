package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsExternalSystemAccessRoleFinding struct {
	AccessibleSystemIds   *[]string              `json:"accessibleSystemIds,omitempty"`
	CreatedDateTime       *string                `json:"createdDateTime,omitempty"`
	Id                    *string                `json:"id,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
	Role                  *AwsRole               `json:"role,omitempty"`
}
