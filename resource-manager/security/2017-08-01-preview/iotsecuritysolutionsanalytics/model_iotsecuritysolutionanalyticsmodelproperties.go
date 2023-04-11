package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionAnalyticsModelProperties struct {
	DevicesMetrics                     *[]IoTSecuritySolutionAnalyticsModelPropertiesDevicesMetricsInlined `json:"devicesMetrics,omitempty"`
	Metrics                            *IoTSeverityMetrics                                                 `json:"metrics,omitempty"`
	MostPrevalentDeviceAlerts          *IoTSecurityDeviceAlertsList                                        `json:"mostPrevalentDeviceAlerts,omitempty"`
	MostPrevalentDeviceRecommendations *IoTSecurityDeviceRecommendationsList                               `json:"mostPrevalentDeviceRecommendations,omitempty"`
	TopAlertedDevices                  *IoTSecurityAlertedDevicesList                                      `json:"topAlertedDevices,omitempty"`
	UnhealthyDeviceCount               *int64                                                              `json:"unhealthyDeviceCount,omitempty"`
}
