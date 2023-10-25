package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceVcoresCapability struct {
	IncludedMaxSize                                               *MaxSizeCapability                                   `json:"includedMaxSize,omitempty"`
	IncludedStorageIOps                                           *int64                                               `json:"includedStorageIOps,omitempty"`
	IncludedStorageThroughputMBps                                 *int64                                               `json:"includedStorageThroughputMBps,omitempty"`
	InstancePoolSupported                                         *bool                                                `json:"instancePoolSupported,omitempty"`
	IopsIncludedValueOverrideFactorPerSelectedStorageGB           *float64                                             `json:"iopsIncludedValueOverrideFactorPerSelectedStorageGB,omitempty"`
	IopsMinValueOverrideFactorPerSelectedStorageGB                *float64                                             `json:"iopsMinValueOverrideFactorPerSelectedStorageGB,omitempty"`
	Name                                                          *string                                              `json:"name,omitempty"`
	Reason                                                        *string                                              `json:"reason,omitempty"`
	StandaloneSupported                                           *bool                                                `json:"standaloneSupported,omitempty"`
	Status                                                        *CapabilityStatus                                    `json:"status,omitempty"`
	SupportedMaintenanceConfigurations                            *[]ManagedInstanceMaintenanceConfigurationCapability `json:"supportedMaintenanceConfigurations,omitempty"`
	SupportedStorageIOps                                          *MaxLimitRangeCapability                             `json:"supportedStorageIOps,omitempty"`
	SupportedStorageSizes                                         *[]MaxSizeRangeCapability                            `json:"supportedStorageSizes,omitempty"`
	SupportedStorageThroughputMBps                                *MaxLimitRangeCapability                             `json:"supportedStorageThroughputMBps,omitempty"`
	ThroughputMBpsIncludedValueOverrideFactorPerSelectedStorageGB *float64                                             `json:"throughputMBpsIncludedValueOverrideFactorPerSelectedStorageGB,omitempty"`
	ThroughputMBpsMinValueOverrideFactorPerSelectedStorageGB      *float64                                             `json:"throughputMBpsMinValueOverrideFactorPerSelectedStorageGB,omitempty"`
	Value                                                         *int64                                               `json:"value,omitempty"`
}
