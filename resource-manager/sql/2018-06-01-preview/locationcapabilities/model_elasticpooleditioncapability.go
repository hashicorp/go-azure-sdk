package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolEditionCapability struct {
	Name                                  *string                                  `json:"name,omitempty"`
	Reason                                *string                                  `json:"reason,omitempty"`
	Status                                *CapabilityStatus                        `json:"status,omitempty"`
	SupportedElasticPoolPerformanceLevels *[]ElasticPoolPerformanceLevelCapability `json:"supportedElasticPoolPerformanceLevels,omitempty"`
	ZoneRedundant                         *bool                                    `json:"zoneRedundant,omitempty"`
}
