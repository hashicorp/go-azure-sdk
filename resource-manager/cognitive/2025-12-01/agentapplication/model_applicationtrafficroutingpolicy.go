package agentapplication

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTrafficRoutingPolicy struct {
	Protocol *TrafficRoutingProtocol `json:"protocol,omitempty"`
	Rules    *[]TrafficRoutingRule   `json:"rules,omitempty"`
}
