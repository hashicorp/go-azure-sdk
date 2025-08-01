package getoperationresult

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringSettings struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettings `json:"azureMonitorAlertSettings,omitempty"`
	ClassicAlertSettings      *ClassicAlertSettings      `json:"classicAlertSettings,omitempty"`
}
