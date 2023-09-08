package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedundantAssignmentAlertConfiguration struct {
	AlertDefinition   *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`
	AlertDefinitionId *string                               `json:"alertDefinitionId,omitempty"`
	Duration          *string                               `json:"duration,omitempty"`
	Id                *string                               `json:"id,omitempty"`
	IsEnabled         *bool                                 `json:"isEnabled,omitempty"`
	ODataType         *string                               `json:"@odata.type,omitempty"`
	ScopeId           *string                               `json:"scopeId,omitempty"`
	ScopeType         *string                               `json:"scopeType,omitempty"`
}
