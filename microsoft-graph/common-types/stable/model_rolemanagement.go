package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagement struct {
	Directory             *RbacApplication `json:"directory,omitempty"`
	EntitlementManagement *RbacApplication `json:"entitlementManagement,omitempty"`
	ODataType             *string          `json:"@odata.type,omitempty"`
}
