package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSetupInfo struct {
	DefaultRolesSettings  *PrivilegedRoleSettings     `json:"defaultRolesSettings,omitempty"`
	FirstTimeSetup        *bool                       `json:"firstTimeSetup,omitempty"`
	Id                    *string                     `json:"id,omitempty"`
	ODataType             *string                     `json:"@odata.type,omitempty"`
	RelevantRolesSettings *[]string                   `json:"relevantRolesSettings,omitempty"`
	SetupStatus           *TenantSetupInfoSetupStatus `json:"setupStatus,omitempty"`
	SkipSetup             *bool                       `json:"skipSetup,omitempty"`
	UserRolesActions      *string                     `json:"userRolesActions,omitempty"`
}
