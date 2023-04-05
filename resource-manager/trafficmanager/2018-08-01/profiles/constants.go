package profiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedEndpointRecordType string

const (
	AllowedEndpointRecordTypeAny            AllowedEndpointRecordType = "Any"
	AllowedEndpointRecordTypeDomainName     AllowedEndpointRecordType = "DomainName"
	AllowedEndpointRecordTypeIPvFourAddress AllowedEndpointRecordType = "IPv4Address"
	AllowedEndpointRecordTypeIPvSixAddress  AllowedEndpointRecordType = "IPv6Address"
)

func PossibleValuesForAllowedEndpointRecordType() []string {
	return []string{
		string(AllowedEndpointRecordTypeAny),
		string(AllowedEndpointRecordTypeDomainName),
		string(AllowedEndpointRecordTypeIPvFourAddress),
		string(AllowedEndpointRecordTypeIPvSixAddress),
	}
}

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

type MonitorProtocol string

const (
	MonitorProtocolHTTP  MonitorProtocol = "HTTP"
	MonitorProtocolHTTPS MonitorProtocol = "HTTPS"
	MonitorProtocolTCP   MonitorProtocol = "TCP"
)

func PossibleValuesForMonitorProtocol() []string {
	return []string{
		string(MonitorProtocolHTTP),
		string(MonitorProtocolHTTPS),
		string(MonitorProtocolTCP),
	}
}

type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = "CheckingEndpoints"
	ProfileMonitorStatusDegraded          ProfileMonitorStatus = "Degraded"
	ProfileMonitorStatusDisabled          ProfileMonitorStatus = "Disabled"
	ProfileMonitorStatusInactive          ProfileMonitorStatus = "Inactive"
	ProfileMonitorStatusOnline            ProfileMonitorStatus = "Online"
)

func PossibleValuesForProfileMonitorStatus() []string {
	return []string{
		string(ProfileMonitorStatusCheckingEndpoints),
		string(ProfileMonitorStatusDegraded),
		string(ProfileMonitorStatusDisabled),
		string(ProfileMonitorStatusInactive),
		string(ProfileMonitorStatusOnline),
	}
}

type ProfileStatus string

const (
	ProfileStatusDisabled ProfileStatus = "Disabled"
	ProfileStatusEnabled  ProfileStatus = "Enabled"
)

func PossibleValuesForProfileStatus() []string {
	return []string{
		string(ProfileStatusDisabled),
		string(ProfileStatusEnabled),
	}
}

type TrafficRoutingMethod string

const (
	TrafficRoutingMethodGeographic  TrafficRoutingMethod = "Geographic"
	TrafficRoutingMethodMultiValue  TrafficRoutingMethod = "MultiValue"
	TrafficRoutingMethodPerformance TrafficRoutingMethod = "Performance"
	TrafficRoutingMethodPriority    TrafficRoutingMethod = "Priority"
	TrafficRoutingMethodSubnet      TrafficRoutingMethod = "Subnet"
	TrafficRoutingMethodWeighted    TrafficRoutingMethod = "Weighted"
)

func PossibleValuesForTrafficRoutingMethod() []string {
	return []string{
		string(TrafficRoutingMethodGeographic),
		string(TrafficRoutingMethodMultiValue),
		string(TrafficRoutingMethodPerformance),
		string(TrafficRoutingMethodPriority),
		string(TrafficRoutingMethodSubnet),
		string(TrafficRoutingMethodWeighted),
	}
}

type TrafficViewEnrollmentStatus string

const (
	TrafficViewEnrollmentStatusDisabled TrafficViewEnrollmentStatus = "Disabled"
	TrafficViewEnrollmentStatusEnabled  TrafficViewEnrollmentStatus = "Enabled"
)

func PossibleValuesForTrafficViewEnrollmentStatus() []string {
	return []string{
		string(TrafficViewEnrollmentStatusDisabled),
		string(TrafficViewEnrollmentStatusEnabled),
	}
}
