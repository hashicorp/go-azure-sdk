package discoveredassetendpointprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredAssetEndpointProfileUpdateProperties struct {
	AdditionalConfiguration        *string                 `json:"additionalConfiguration,omitempty"`
	DiscoveryId                    *string                 `json:"discoveryId,omitempty"`
	EndpointProfileType            *string                 `json:"endpointProfileType,omitempty"`
	SupportedAuthenticationMethods *[]AuthenticationMethod `json:"supportedAuthenticationMethods,omitempty"`
	TargetAddress                  *string                 `json:"targetAddress,omitempty"`
	Version                        *int64                  `json:"version,omitempty"`
}
