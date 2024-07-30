package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreAuthorizedApplication struct {
	AppId         *string   `json:"appId,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
	PermissionIds *[]string `json:"permissionIds,omitempty"`
}
