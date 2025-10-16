package namespacediscovereddevices

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredInboundEndpointsUpdate struct {
	AdditionalConfiguration        *string                 `json:"additionalConfiguration,omitempty"`
	Address                        *string                 `json:"address,omitempty"`
	EndpointType                   *string                 `json:"endpointType,omitempty"`
	LastUpdatedOn                  *string                 `json:"lastUpdatedOn,omitempty"`
	SupportedAuthenticationMethods *[]AuthenticationMethod `json:"supportedAuthenticationMethods,omitempty"`
	Version                        *string                 `json:"version,omitempty"`
}

func (o *DiscoveredInboundEndpointsUpdate) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DiscoveredInboundEndpointsUpdate) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}
