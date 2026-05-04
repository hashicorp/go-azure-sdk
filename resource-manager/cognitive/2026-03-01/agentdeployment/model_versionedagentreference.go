package agentdeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VersionedAgentReference struct {
	AgentId      *string `json:"agentId,omitempty"`
	AgentName    *string `json:"agentName,omitempty"`
	AgentVersion *string `json:"agentVersion,omitempty"`
}
