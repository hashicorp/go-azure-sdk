package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentExtensionDefinition struct {
	Alias    *string                                   `json:"alias,omitempty"`
	Config   *map[string]DeploymentExtensionConfigItem `json:"config,omitempty"`
	ConfigId *string                                   `json:"configId,omitempty"`
	Name     *string                                   `json:"name,omitempty"`
	Version  *string                                   `json:"version,omitempty"`
}
