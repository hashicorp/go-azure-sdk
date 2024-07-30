package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedPermissionClassification struct {
	Classification *DelegatedPermissionClassificationClassification `json:"classification,omitempty"`
	Id             *string                                          `json:"id,omitempty"`
	ODataType      *string                                          `json:"@odata.type,omitempty"`
	PermissionId   *string                                          `json:"permissionId,omitempty"`
	PermissionName *string                                          `json:"permissionName,omitempty"`
}
