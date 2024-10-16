package discoveredassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredAssetProperties struct {
	AssetEndpointProfileRef      string               `json:"assetEndpointProfileRef"`
	Datasets                     *[]DiscoveredDataset `json:"datasets,omitempty"`
	DefaultDatasetsConfiguration *string              `json:"defaultDatasetsConfiguration,omitempty"`
	DefaultEventsConfiguration   *string              `json:"defaultEventsConfiguration,omitempty"`
	DefaultTopic                 *Topic               `json:"defaultTopic,omitempty"`
	DiscoveryId                  string               `json:"discoveryId"`
	DocumentationUri             *string              `json:"documentationUri,omitempty"`
	Events                       *[]DiscoveredEvent   `json:"events,omitempty"`
	HardwareRevision             *string              `json:"hardwareRevision,omitempty"`
	Manufacturer                 *string              `json:"manufacturer,omitempty"`
	ManufacturerUri              *string              `json:"manufacturerUri,omitempty"`
	Model                        *string              `json:"model,omitempty"`
	ProductCode                  *string              `json:"productCode,omitempty"`
	ProvisioningState            *ProvisioningState   `json:"provisioningState,omitempty"`
	SerialNumber                 *string              `json:"serialNumber,omitempty"`
	SoftwareRevision             *string              `json:"softwareRevision,omitempty"`
	Version                      int64                `json:"version"`
}
