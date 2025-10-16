package namespaceassets

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceAssetProperties struct {
	AssetTypeRefs                        *[]string               `json:"assetTypeRefs,omitempty"`
	Attributes                           *map[string]interface{} `json:"attributes,omitempty"`
	Datasets                             *[]NamespaceDataset     `json:"datasets,omitempty"`
	DefaultDatasetsConfiguration         *string                 `json:"defaultDatasetsConfiguration,omitempty"`
	DefaultDatasetsDestinations          *[]DatasetDestination   `json:"defaultDatasetsDestinations,omitempty"`
	DefaultEventsConfiguration           *string                 `json:"defaultEventsConfiguration,omitempty"`
	DefaultEventsDestinations            *[]EventDestination     `json:"defaultEventsDestinations,omitempty"`
	DefaultManagementGroupsConfiguration *string                 `json:"defaultManagementGroupsConfiguration,omitempty"`
	DefaultStreamsConfiguration          *string                 `json:"defaultStreamsConfiguration,omitempty"`
	DefaultStreamsDestinations           *[]StreamDestination    `json:"defaultStreamsDestinations,omitempty"`
	Description                          *string                 `json:"description,omitempty"`
	DeviceRef                            DeviceRef               `json:"deviceRef"`
	DiscoveredAssetRefs                  *[]string               `json:"discoveredAssetRefs,omitempty"`
	DisplayName                          *string                 `json:"displayName,omitempty"`
	DocumentationUri                     *string                 `json:"documentationUri,omitempty"`
	Enabled                              *bool                   `json:"enabled,omitempty"`
	EventGroups                          *[]NamespaceEventGroup  `json:"eventGroups,omitempty"`
	ExternalAssetId                      *string                 `json:"externalAssetId,omitempty"`
	HardwareRevision                     *string                 `json:"hardwareRevision,omitempty"`
	LastTransitionTime                   *string                 `json:"lastTransitionTime,omitempty"`
	ManagementGroups                     *[]ManagementGroup      `json:"managementGroups,omitempty"`
	Manufacturer                         *string                 `json:"manufacturer,omitempty"`
	ManufacturerUri                      *string                 `json:"manufacturerUri,omitempty"`
	Model                                *string                 `json:"model,omitempty"`
	ProductCode                          *string                 `json:"productCode,omitempty"`
	ProvisioningState                    *ProvisioningState      `json:"provisioningState,omitempty"`
	SerialNumber                         *string                 `json:"serialNumber,omitempty"`
	SoftwareRevision                     *string                 `json:"softwareRevision,omitempty"`
	Status                               *NamespaceAssetStatus   `json:"status,omitempty"`
	Streams                              *[]NamespaceStream      `json:"streams,omitempty"`
	Uuid                                 *string                 `json:"uuid,omitempty"`
	Version                              *int64                  `json:"version,omitempty"`
}

func (o *NamespaceAssetProperties) GetLastTransitionTimeAsTime() (*time.Time, error) {
	if o.LastTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceAssetProperties) SetLastTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastTransitionTime = &formatted
}

var _ json.Unmarshaler = &NamespaceAssetProperties{}

func (s *NamespaceAssetProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssetTypeRefs                        *[]string               `json:"assetTypeRefs,omitempty"`
		Attributes                           *map[string]interface{} `json:"attributes,omitempty"`
		Datasets                             *[]NamespaceDataset     `json:"datasets,omitempty"`
		DefaultDatasetsConfiguration         *string                 `json:"defaultDatasetsConfiguration,omitempty"`
		DefaultEventsConfiguration           *string                 `json:"defaultEventsConfiguration,omitempty"`
		DefaultManagementGroupsConfiguration *string                 `json:"defaultManagementGroupsConfiguration,omitempty"`
		DefaultStreamsConfiguration          *string                 `json:"defaultStreamsConfiguration,omitempty"`
		Description                          *string                 `json:"description,omitempty"`
		DeviceRef                            DeviceRef               `json:"deviceRef"`
		DiscoveredAssetRefs                  *[]string               `json:"discoveredAssetRefs,omitempty"`
		DisplayName                          *string                 `json:"displayName,omitempty"`
		DocumentationUri                     *string                 `json:"documentationUri,omitempty"`
		Enabled                              *bool                   `json:"enabled,omitempty"`
		EventGroups                          *[]NamespaceEventGroup  `json:"eventGroups,omitempty"`
		ExternalAssetId                      *string                 `json:"externalAssetId,omitempty"`
		HardwareRevision                     *string                 `json:"hardwareRevision,omitempty"`
		LastTransitionTime                   *string                 `json:"lastTransitionTime,omitempty"`
		ManagementGroups                     *[]ManagementGroup      `json:"managementGroups,omitempty"`
		Manufacturer                         *string                 `json:"manufacturer,omitempty"`
		ManufacturerUri                      *string                 `json:"manufacturerUri,omitempty"`
		Model                                *string                 `json:"model,omitempty"`
		ProductCode                          *string                 `json:"productCode,omitempty"`
		ProvisioningState                    *ProvisioningState      `json:"provisioningState,omitempty"`
		SerialNumber                         *string                 `json:"serialNumber,omitempty"`
		SoftwareRevision                     *string                 `json:"softwareRevision,omitempty"`
		Status                               *NamespaceAssetStatus   `json:"status,omitempty"`
		Streams                              *[]NamespaceStream      `json:"streams,omitempty"`
		Uuid                                 *string                 `json:"uuid,omitempty"`
		Version                              *int64                  `json:"version,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssetTypeRefs = decoded.AssetTypeRefs
	s.Attributes = decoded.Attributes
	s.Datasets = decoded.Datasets
	s.DefaultDatasetsConfiguration = decoded.DefaultDatasetsConfiguration
	s.DefaultEventsConfiguration = decoded.DefaultEventsConfiguration
	s.DefaultManagementGroupsConfiguration = decoded.DefaultManagementGroupsConfiguration
	s.DefaultStreamsConfiguration = decoded.DefaultStreamsConfiguration
	s.Description = decoded.Description
	s.DeviceRef = decoded.DeviceRef
	s.DiscoveredAssetRefs = decoded.DiscoveredAssetRefs
	s.DisplayName = decoded.DisplayName
	s.DocumentationUri = decoded.DocumentationUri
	s.Enabled = decoded.Enabled
	s.EventGroups = decoded.EventGroups
	s.ExternalAssetId = decoded.ExternalAssetId
	s.HardwareRevision = decoded.HardwareRevision
	s.LastTransitionTime = decoded.LastTransitionTime
	s.ManagementGroups = decoded.ManagementGroups
	s.Manufacturer = decoded.Manufacturer
	s.ManufacturerUri = decoded.ManufacturerUri
	s.Model = decoded.Model
	s.ProductCode = decoded.ProductCode
	s.ProvisioningState = decoded.ProvisioningState
	s.SerialNumber = decoded.SerialNumber
	s.SoftwareRevision = decoded.SoftwareRevision
	s.Status = decoded.Status
	s.Streams = decoded.Streams
	s.Uuid = decoded.Uuid
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceAssetProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultDatasetsDestinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DefaultDatasetsDestinations into list []json.RawMessage: %+v", err)
		}

		output := make([]DatasetDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDatasetDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DefaultDatasetsDestinations' for 'NamespaceAssetProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DefaultDatasetsDestinations = &output
	}

	if v, ok := temp["defaultEventsDestinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DefaultEventsDestinations into list []json.RawMessage: %+v", err)
		}

		output := make([]EventDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEventDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DefaultEventsDestinations' for 'NamespaceAssetProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DefaultEventsDestinations = &output
	}

	if v, ok := temp["defaultStreamsDestinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DefaultStreamsDestinations into list []json.RawMessage: %+v", err)
		}

		output := make([]StreamDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalStreamDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DefaultStreamsDestinations' for 'NamespaceAssetProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DefaultStreamsDestinations = &output
	}

	return nil
}
