package connectionmonitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionMonitorResultProperties struct {
	AutoStart                   *bool                                 `json:"autoStart,omitempty"`
	ConnectionMonitorType       *ConnectionMonitorType                `json:"connectionMonitorType,omitempty"`
	Destination                 *ConnectionMonitorDestination         `json:"destination,omitempty"`
	Endpoints                   *[]ConnectionMonitorEndpoint          `json:"endpoints,omitempty"`
	MonitoringIntervalInSeconds *int64                                `json:"monitoringIntervalInSeconds,omitempty"`
	MonitoringStatus            *string                               `json:"monitoringStatus,omitempty"`
	Notes                       *string                               `json:"notes,omitempty"`
	Outputs                     *[]ConnectionMonitorOutput            `json:"outputs,omitempty"`
	ProvisioningState           *ProvisioningState                    `json:"provisioningState,omitempty"`
	Source                      *ConnectionMonitorSource              `json:"source,omitempty"`
	StartTime                   *string                               `json:"startTime,omitempty"`
	TestConfigurations          *[]ConnectionMonitorTestConfiguration `json:"testConfigurations,omitempty"`
	TestGroups                  *[]ConnectionMonitorTestGroup         `json:"testGroups,omitempty"`
}
