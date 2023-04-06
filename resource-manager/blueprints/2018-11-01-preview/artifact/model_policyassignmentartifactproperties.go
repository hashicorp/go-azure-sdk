package artifact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentArtifactProperties struct {
	DependsOn          *[]string                 `json:"dependsOn,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	DisplayName        *string                   `json:"displayName,omitempty"`
	Parameters         map[string]ParameterValue `json:"parameters"`
	PolicyDefinitionId string                    `json:"policyDefinitionId"`
	ResourceGroup      *string                   `json:"resourceGroup,omitempty"`
}
