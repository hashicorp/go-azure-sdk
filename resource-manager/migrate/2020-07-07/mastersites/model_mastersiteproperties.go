package mastersites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MasterSiteProperties struct {
	AllowMultipleSites          *bool                        `json:"allowMultipleSites,omitempty"`
	CustomerStorageAccountArmId *string                      `json:"customerStorageAccountArmId,omitempty"`
	PrivateEndpointConnections  *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess         *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
	Sites                       *[]string                    `json:"sites,omitempty"`
}
