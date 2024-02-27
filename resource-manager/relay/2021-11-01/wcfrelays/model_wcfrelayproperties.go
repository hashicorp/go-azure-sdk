package wcfrelays

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *WcfRelayProperties) GetCreatedAtAsTime() (*time.Time, error) {
	if o.CreatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *WcfRelayProperties) GetUpdatedAtAsTime() (*time.Time, error) {
	if o.UpdatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpdatedAt, "2006-01-02T15:04:05Z07:00")
}
