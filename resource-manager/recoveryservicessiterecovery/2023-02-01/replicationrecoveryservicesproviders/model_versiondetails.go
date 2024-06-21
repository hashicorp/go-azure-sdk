package replicationrecoveryservicesproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VersionDetails struct {
	ExpiryDate *string             `json:"expiryDate,omitempty"`
	Status     *AgentVersionStatus `json:"status,omitempty"`
	Version    *string             `json:"version,omitempty"`
}
