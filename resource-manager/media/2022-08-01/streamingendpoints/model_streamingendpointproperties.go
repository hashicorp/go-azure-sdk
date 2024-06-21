package streamingendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingEndpointProperties struct {
	AccessControl           *StreamingEndpointAccessControl `json:"accessControl,omitempty"`
	AvailabilitySetName     *string                         `json:"availabilitySetName,omitempty"`
	CdnEnabled              *bool                           `json:"cdnEnabled,omitempty"`
	CdnProfile              *string                         `json:"cdnProfile,omitempty"`
	CdnProvider             *string                         `json:"cdnProvider,omitempty"`
	Created                 *string                         `json:"created,omitempty"`
	CrossSiteAccessPolicies *CrossSiteAccessPolicies        `json:"crossSiteAccessPolicies,omitempty"`
	CustomHostNames         *[]string                       `json:"customHostNames,omitempty"`
	Description             *string                         `json:"description,omitempty"`
	FreeTrialEndTime        *string                         `json:"freeTrialEndTime,omitempty"`
	HostName                *string                         `json:"hostName,omitempty"`
	LastModified            *string                         `json:"lastModified,omitempty"`
	MaxCacheAge             *int64                          `json:"maxCacheAge,omitempty"`
	ProvisioningState       *string                         `json:"provisioningState,omitempty"`
	ResourceState           *StreamingEndpointResourceState `json:"resourceState,omitempty"`
	ScaleUnits              int64                           `json:"scaleUnits"`
}
