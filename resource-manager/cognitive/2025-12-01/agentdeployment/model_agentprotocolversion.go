package agentdeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentProtocolVersion struct {
	Protocol *AgentProtocol `json:"protocol,omitempty"`
	Version  *string        `json:"version,omitempty"`
}
