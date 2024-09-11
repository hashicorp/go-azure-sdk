package origingroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OriginGroupProperties struct {
	HealthProbeSettings                                   *HealthProbeParameters                       `json:"healthProbeSettings,omitempty"`
	Origins                                               *[]ResourceReference                         `json:"origins,omitempty"`
	ProvisioningState                                     *OriginGroupProvisioningState                `json:"provisioningState,omitempty"`
	ResourceState                                         *OriginGroupResourceState                    `json:"resourceState,omitempty"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParameters `json:"responseBasedOriginErrorDetectionSettings,omitempty"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int64                                       `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
}
