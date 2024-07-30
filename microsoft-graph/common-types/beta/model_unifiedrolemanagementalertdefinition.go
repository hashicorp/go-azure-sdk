package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertDefinition struct {
	Description     *string                                            `json:"description,omitempty"`
	DisplayName     *string                                            `json:"displayName,omitempty"`
	HowToPrevent    *string                                            `json:"howToPrevent,omitempty"`
	Id              *string                                            `json:"id,omitempty"`
	IsConfigurable  *bool                                              `json:"isConfigurable,omitempty"`
	IsRemediatable  *bool                                              `json:"isRemediatable,omitempty"`
	MitigationSteps *string                                            `json:"mitigationSteps,omitempty"`
	ODataType       *string                                            `json:"@odata.type,omitempty"`
	ScopeId         *string                                            `json:"scopeId,omitempty"`
	ScopeType       *string                                            `json:"scopeType,omitempty"`
	SecurityImpact  *string                                            `json:"securityImpact,omitempty"`
	SeverityLevel   *UnifiedRoleManagementAlertDefinitionSeverityLevel `json:"severityLevel,omitempty"`
}
