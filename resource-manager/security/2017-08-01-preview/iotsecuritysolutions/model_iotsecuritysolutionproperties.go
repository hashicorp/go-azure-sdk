package iotsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionProperties struct {
	AutoDiscoveredResources      *[]string                                `json:"autoDiscoveredResources,omitempty"`
	DisabledDataSources          *[]DataSource                            `json:"disabledDataSources,omitempty"`
	DisplayName                  string                                   `json:"displayName"`
	Export                       *[]ExportData                            `json:"export,omitempty"`
	IotHubs                      []string                                 `json:"iotHubs"`
	RecommendationsConfiguration *[]RecommendationConfigurationProperties `json:"recommendationsConfiguration,omitempty"`
	Status                       *SecuritySolutionStatus                  `json:"status,omitempty"`
	UserDefinedResources         *UserDefinedResourcesProperties          `json:"userDefinedResources,omitempty"`
	Workspace                    string                                   `json:"workspace"`
}
