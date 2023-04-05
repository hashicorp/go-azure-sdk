package performconnectivitycheck

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionStatus string

const (
	ConnectionStatusConnected    ConnectionStatus = "Connected"
	ConnectionStatusDegraded     ConnectionStatus = "Degraded"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusUnknown      ConnectionStatus = "Unknown"
)

func PossibleValuesForConnectionStatus() []string {
	return []string{
		string(ConnectionStatusConnected),
		string(ConnectionStatusDegraded),
		string(ConnectionStatusDisconnected),
		string(ConnectionStatusUnknown),
	}
}

type ConnectivityCheckProtocol string

const (
	ConnectivityCheckProtocolHTTP  ConnectivityCheckProtocol = "HTTP"
	ConnectivityCheckProtocolHTTPS ConnectivityCheckProtocol = "HTTPS"
	ConnectivityCheckProtocolTCP   ConnectivityCheckProtocol = "TCP"
)

func PossibleValuesForConnectivityCheckProtocol() []string {
	return []string{
		string(ConnectivityCheckProtocolHTTP),
		string(ConnectivityCheckProtocolHTTPS),
		string(ConnectivityCheckProtocolTCP),
	}
}

type IssueType string

const (
	IssueTypeAgentStopped        IssueType = "AgentStopped"
	IssueTypeDnsResolution       IssueType = "DnsResolution"
	IssueTypeGuestFirewall       IssueType = "GuestFirewall"
	IssueTypeNetworkSecurityRule IssueType = "NetworkSecurityRule"
	IssueTypePlatform            IssueType = "Platform"
	IssueTypePortThrottled       IssueType = "PortThrottled"
	IssueTypeSocketBind          IssueType = "SocketBind"
	IssueTypeUnknown             IssueType = "Unknown"
	IssueTypeUserDefinedRoute    IssueType = "UserDefinedRoute"
)

func PossibleValuesForIssueType() []string {
	return []string{
		string(IssueTypeAgentStopped),
		string(IssueTypeDnsResolution),
		string(IssueTypeGuestFirewall),
		string(IssueTypeNetworkSecurityRule),
		string(IssueTypePlatform),
		string(IssueTypePortThrottled),
		string(IssueTypeSocketBind),
		string(IssueTypeUnknown),
		string(IssueTypeUserDefinedRoute),
	}
}

type Method string

const (
	MethodGET  Method = "GET"
	MethodPOST Method = "POST"
)

func PossibleValuesForMethod() []string {
	return []string{
		string(MethodGET),
		string(MethodPOST),
	}
}

type Origin string

const (
	OriginInbound  Origin = "Inbound"
	OriginLocal    Origin = "Local"
	OriginOutbound Origin = "Outbound"
)

func PossibleValuesForOrigin() []string {
	return []string{
		string(OriginInbound),
		string(OriginLocal),
		string(OriginOutbound),
	}
}

type PreferredIPVersion string

const (
	PreferredIPVersionIPvFour PreferredIPVersion = "IPv4"
)

func PossibleValuesForPreferredIPVersion() []string {
	return []string{
		string(PreferredIPVersionIPvFour),
	}
}

type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityWarning Severity = "Warning"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityError),
		string(SeverityWarning),
	}
}
