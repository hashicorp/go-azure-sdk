package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlert struct {
	AlertConfigurations *[]UnifiedRoleManagementAlertConfiguration `json:"alertConfigurations,omitempty"`
	AlertDefinitions    *[]UnifiedRoleManagementAlertDefinition    `json:"alertDefinitions,omitempty"`
	Alerts              *[]UnifiedRoleManagementAlert              `json:"alerts,omitempty"`
	Id                  *string                                    `json:"id,omitempty"`
	ODataType           *string                                    `json:"@odata.type,omitempty"`
	Operations          *[]LongRunningOperation                    `json:"operations,omitempty"`
}
