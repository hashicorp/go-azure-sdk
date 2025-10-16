package namespacedevices

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDeviceProperties struct {
	Attributes             *map[string]interface{} `json:"attributes,omitempty"`
	DiscoveredDeviceRef    *string                 `json:"discoveredDeviceRef,omitempty"`
	Enabled                *bool                   `json:"enabled,omitempty"`
	Endpoints              *MessagingEndpoints     `json:"endpoints,omitempty"`
	ExternalDeviceId       *string                 `json:"externalDeviceId,omitempty"`
	LastTransitionTime     *string                 `json:"lastTransitionTime,omitempty"`
	Manufacturer           *string                 `json:"manufacturer,omitempty"`
	Model                  *string                 `json:"model,omitempty"`
	OperatingSystem        *string                 `json:"operatingSystem,omitempty"`
	OperatingSystemVersion *string                 `json:"operatingSystemVersion,omitempty"`
	ProvisioningState      *ProvisioningState      `json:"provisioningState,omitempty"`
	Status                 *DeviceStatus           `json:"status,omitempty"`
	Uuid                   *string                 `json:"uuid,omitempty"`
	Version                *int64                  `json:"version,omitempty"`
}

func (o *NamespaceDeviceProperties) GetLastTransitionTimeAsTime() (*time.Time, error) {
	if o.LastTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceDeviceProperties) SetLastTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastTransitionTime = &formatted
}
