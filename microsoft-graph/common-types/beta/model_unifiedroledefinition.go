package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleDefinition struct {
	AllowedPrincipalTypes   *UnifiedRoleDefinitionAllowedPrincipalTypes `json:"allowedPrincipalTypes,omitempty"`
	Description             *string                                     `json:"description,omitempty"`
	DisplayName             *string                                     `json:"displayName,omitempty"`
	Id                      *string                                     `json:"id,omitempty"`
	InheritsPermissionsFrom *[]UnifiedRoleDefinition                    `json:"inheritsPermissionsFrom,omitempty"`
	IsBuiltIn               *bool                                       `json:"isBuiltIn,omitempty"`
	IsEnabled               *bool                                       `json:"isEnabled,omitempty"`
	IsPrivileged            *bool                                       `json:"isPrivileged,omitempty"`
	ODataType               *string                                     `json:"@odata.type,omitempty"`
	ResourceScopes          *[]string                                   `json:"resourceScopes,omitempty"`
	RolePermissions         *[]UnifiedRolePermission                    `json:"rolePermissions,omitempty"`
	TemplateId              *string                                     `json:"templateId,omitempty"`
	Version                 *string                                     `json:"version,omitempty"`
}
