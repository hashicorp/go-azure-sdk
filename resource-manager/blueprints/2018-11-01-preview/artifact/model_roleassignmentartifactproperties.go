package artifact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentArtifactProperties struct {
	DependsOn        *[]string   `json:"dependsOn,omitempty"`
	Description      *string     `json:"description,omitempty"`
	DisplayName      *string     `json:"displayName,omitempty"`
	PrincipalIds     interface{} `json:"principalIds"`
	ResourceGroup    *string     `json:"resourceGroup,omitempty"`
	RoleDefinitionId string      `json:"roleDefinitionId"`
}
