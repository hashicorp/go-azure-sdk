package agentregistrationinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentRegistrationRegenerateKeyParameter struct {
	KeyName  AgentRegistrationKeyName `json:"keyName"`
	Location *string                  `json:"location,omitempty"`
	Name     *string                  `json:"name,omitempty"`
	Tags     *map[string]string       `json:"tags,omitempty"`
}
