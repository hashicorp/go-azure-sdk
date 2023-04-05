package endpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint EndpointMonitorStatus = "CheckingEndpoint"
	EndpointMonitorStatusDegraded         EndpointMonitorStatus = "Degraded"
	EndpointMonitorStatusDisabled         EndpointMonitorStatus = "Disabled"
	EndpointMonitorStatusInactive         EndpointMonitorStatus = "Inactive"
	EndpointMonitorStatusOnline           EndpointMonitorStatus = "Online"
	EndpointMonitorStatusStopped          EndpointMonitorStatus = "Stopped"
)

func PossibleValuesForEndpointMonitorStatus() []string {
	return []string{
		string(EndpointMonitorStatusCheckingEndpoint),
		string(EndpointMonitorStatusDegraded),
		string(EndpointMonitorStatusDisabled),
		string(EndpointMonitorStatusInactive),
		string(EndpointMonitorStatusOnline),
		string(EndpointMonitorStatusStopped),
	}
}

type EndpointStatus string

const (
	EndpointStatusDisabled EndpointStatus = "Disabled"
	EndpointStatusEnabled  EndpointStatus = "Enabled"
)

func PossibleValuesForEndpointStatus() []string {
	return []string{
		string(EndpointStatusDisabled),
		string(EndpointStatusEnabled),
	}
}

type EndpointType string

const (
	EndpointTypeAzureEndpoints    EndpointType = "AzureEndpoints"
	EndpointTypeExternalEndpoints EndpointType = "ExternalEndpoints"
	EndpointTypeNestedEndpoints   EndpointType = "NestedEndpoints"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeAzureEndpoints),
		string(EndpointTypeExternalEndpoints),
		string(EndpointTypeNestedEndpoints),
	}
}
