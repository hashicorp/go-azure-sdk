package allowedconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedResource struct {
	ConnectedResourceId *string `json:"connectedResourceId,omitempty"`
	TcpPorts            *string `json:"tcpPorts,omitempty"`
	UdpPorts            *string `json:"udpPorts,omitempty"`
}
