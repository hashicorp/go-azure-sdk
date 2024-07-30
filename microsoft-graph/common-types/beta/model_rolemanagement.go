package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagement struct {
	CloudPC               *RbacApplicationMultiple `json:"cloudPC,omitempty"`
	DeviceManagement      *RbacApplicationMultiple `json:"deviceManagement,omitempty"`
	Directory             *RbacApplication         `json:"directory,omitempty"`
	EnterpriseApps        *[]RbacApplication       `json:"enterpriseApps,omitempty"`
	EntitlementManagement *RbacApplication         `json:"entitlementManagement,omitempty"`
	Exchange              *UnifiedRbacApplication  `json:"exchange,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
}
