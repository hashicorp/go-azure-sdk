package globalreachconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalReachConnectionProperties struct {
	AddressPrefix           *string                                 `json:"addressPrefix,omitempty"`
	AuthorizationKey        *string                                 `json:"authorizationKey,omitempty"`
	CircuitConnectionStatus *GlobalReachConnectionStatus            `json:"circuitConnectionStatus,omitempty"`
	ExpressRouteId          *string                                 `json:"expressRouteId,omitempty"`
	PeerExpressRouteCircuit *string                                 `json:"peerExpressRouteCircuit,omitempty"`
	ProvisioningState       *GlobalReachConnectionProvisioningState `json:"provisioningState,omitempty"`
}
