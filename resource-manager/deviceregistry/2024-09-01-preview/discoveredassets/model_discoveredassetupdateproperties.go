package discoveredassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredAssetUpdateProperties struct {
	Datasets                     *[]DiscoveredDataset `json:"datasets,omitempty"`
	DefaultDatasetsConfiguration *string              `json:"defaultDatasetsConfiguration,omitempty"`
	DefaultEventsConfiguration   *string              `json:"defaultEventsConfiguration,omitempty"`
	DefaultTopic                 *TopicUpdate         `json:"defaultTopic,omitempty"`
	DiscoveryId                  *string              `json:"discoveryId,omitempty"`
	DocumentationUri             *string              `json:"documentationUri,omitempty"`
	Events                       *[]DiscoveredEvent   `json:"events,omitempty"`
	HardwareRevision             *string              `json:"hardwareRevision,omitempty"`
	Manufacturer                 *string              `json:"manufacturer,omitempty"`
	ManufacturerUri              *string              `json:"manufacturerUri,omitempty"`
	Model                        *string              `json:"model,omitempty"`
	ProductCode                  *string              `json:"productCode,omitempty"`
	SerialNumber                 *string              `json:"serialNumber,omitempty"`
	SoftwareRevision             *string              `json:"softwareRevision,omitempty"`
	Version                      *int64               `json:"version,omitempty"`
}
