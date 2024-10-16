package discoveredassetendpointprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredAssetEndpointProfileProperties struct {
	AdditionalConfiguration        *string                 `json:"additionalConfiguration,omitempty"`
	DiscoveryId                    string                  `json:"discoveryId"`
	EndpointProfileType            string                  `json:"endpointProfileType"`
	ProvisioningState              *ProvisioningState      `json:"provisioningState,omitempty"`
	SupportedAuthenticationMethods *[]AuthenticationMethod `json:"supportedAuthenticationMethods,omitempty"`
	TargetAddress                  string                  `json:"targetAddress"`
	Version                        int64                   `json:"version"`
}
