package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCScopedPermission struct {
	ODataType  *string   `json:"@odata.type,omitempty"`
	Permission *string   `json:"permission,omitempty"`
	ScopeIds   *[]string `json:"scopeIds,omitempty"`
}
