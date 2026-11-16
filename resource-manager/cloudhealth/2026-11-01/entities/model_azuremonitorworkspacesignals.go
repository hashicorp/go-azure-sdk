package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureMonitorWorkspaceSignals struct {
	AuthenticationSetting           string                     `json:"authenticationSetting"`
	AzureMonitorWorkspaceResourceId string                     `json:"azureMonitorWorkspaceResourceId"`
	Signals                         *[]PrometheusMetricsSignal `json:"signals,omitempty"`
}
