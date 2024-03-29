package mhsmprivateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateLinkServiceConnectionState struct {
	ActionsRequired *ActionsRequired                        `json:"actionsRequired,omitempty"`
	Description     *string                                 `json:"description,omitempty"`
	Status          *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}
