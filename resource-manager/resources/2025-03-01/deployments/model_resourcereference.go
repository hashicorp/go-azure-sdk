package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceReference struct {
	ApiVersion   *string                        `json:"apiVersion,omitempty"`
	Extension    *DeploymentExtensionDefinition `json:"extension,omitempty"`
	Id           *string                        `json:"id,omitempty"`
	Identifiers  *interface{}                   `json:"identifiers,omitempty"`
	ResourceType *string                        `json:"resourceType,omitempty"`
}
