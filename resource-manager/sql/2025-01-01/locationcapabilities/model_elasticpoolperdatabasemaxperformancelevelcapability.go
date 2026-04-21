package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolPerDatabaseMaxPerformanceLevelCapability struct {
	Limit                                    *float64                                               `json:"limit,omitempty"`
	Reason                                   *string                                                `json:"reason,omitempty"`
	Status                                   *CapabilityStatus                                      `json:"status,omitempty"`
	SupportedPerDatabaseMinPerformanceLevels *[]ElasticPoolPerDatabaseMinPerformanceLevelCapability `json:"supportedPerDatabaseMinPerformanceLevels,omitempty"`
	Unit                                     *PerformanceLevelUnit                                  `json:"unit,omitempty"`
}
