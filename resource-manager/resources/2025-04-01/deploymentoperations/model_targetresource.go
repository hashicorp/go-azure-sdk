package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetResource struct {
	ApiVersion   *string                        `json:"apiVersion,omitempty"`
	Extension    *DeploymentExtensionDefinition `json:"extension,omitempty"`
	Id           *string                        `json:"id,omitempty"`
	Identifiers  *interface{}                   `json:"identifiers,omitempty"`
	ResourceName *string                        `json:"resourceName,omitempty"`
	ResourceType *string                        `json:"resourceType,omitempty"`
	SymbolicName *string                        `json:"symbolicName,omitempty"`
}
