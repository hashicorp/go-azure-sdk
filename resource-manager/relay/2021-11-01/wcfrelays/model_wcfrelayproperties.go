package wcfrelays

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WcfRelayProperties struct {
	CreatedAt                   *string    `json:"createdAt,omitempty"`
	IsDynamic                   *bool      `json:"isDynamic,omitempty"`
	ListenerCount               *int64     `json:"listenerCount,omitempty"`
	RelayType                   *Relaytype `json:"relayType,omitempty"`
	RequiresClientAuthorization *bool      `json:"requiresClientAuthorization,omitempty"`
	RequiresTransportSecurity   *bool      `json:"requiresTransportSecurity,omitempty"`
	UpdatedAt                   *string    `json:"updatedAt,omitempty"`
	UserMetadata                *string    `json:"userMetadata,omitempty"`
}
