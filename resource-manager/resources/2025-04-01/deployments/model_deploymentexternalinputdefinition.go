package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentExternalInputDefinition struct {
	Config *interface{} `json:"config,omitempty"`
	Kind   string       `json:"kind"`
}
