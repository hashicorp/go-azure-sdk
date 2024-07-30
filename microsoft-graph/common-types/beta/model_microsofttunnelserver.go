package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServer struct {
	AgentImageDigest         *string                                        `json:"agentImageDigest,omitempty"`
	DisplayName              *string                                        `json:"displayName,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	LastCheckinDateTime      *string                                        `json:"lastCheckinDateTime,omitempty"`
	ODataType                *string                                        `json:"@odata.type,omitempty"`
	ServerImageDigest        *string                                        `json:"serverImageDigest,omitempty"`
	TunnelServerHealthStatus *MicrosoftTunnelServerTunnelServerHealthStatus `json:"tunnelServerHealthStatus,omitempty"`
}
