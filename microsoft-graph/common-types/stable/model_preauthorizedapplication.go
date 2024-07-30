package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreAuthorizedApplication struct {
	AppId                  *string   `json:"appId,omitempty"`
	DelegatedPermissionIds *[]string `json:"delegatedPermissionIds,omitempty"`
	ODataType              *string   `json:"@odata.type,omitempty"`
}
