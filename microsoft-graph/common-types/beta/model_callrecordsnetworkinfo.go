package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfo struct {
	BandwidthLowEventRatio     *float64                                        `json:"bandwidthLowEventRatio,omitempty"`
	BasicServiceSetIdentifier  *string                                         `json:"basicServiceSetIdentifier,omitempty"`
	ConnectionType             *CallRecordsNetworkInfoConnectionType           `json:"connectionType,omitempty"`
	DelayEventRatio            *float64                                        `json:"delayEventRatio,omitempty"`
	DnsSuffix                  *string                                         `json:"dnsSuffix,omitempty"`
	IpAddress                  *string                                         `json:"ipAddress,omitempty"`
	LinkSpeed                  *int64                                          `json:"linkSpeed,omitempty"`
	MacAddress                 *string                                         `json:"macAddress,omitempty"`
	NetworkTransportProtocol   *CallRecordsNetworkInfoNetworkTransportProtocol `json:"networkTransportProtocol,omitempty"`
	ODataType                  *string                                         `json:"@odata.type,omitempty"`
	Port                       *int64                                          `json:"port,omitempty"`
	ReceivedQualityEventRatio  *float64                                        `json:"receivedQualityEventRatio,omitempty"`
	ReflexiveIPAddress         *string                                         `json:"reflexiveIPAddress,omitempty"`
	RelayIPAddress             *string                                         `json:"relayIPAddress,omitempty"`
	RelayPort                  *int64                                          `json:"relayPort,omitempty"`
	SentQualityEventRatio      *float64                                        `json:"sentQualityEventRatio,omitempty"`
	Subnet                     *string                                         `json:"subnet,omitempty"`
	TraceRouteHops             *[]CallRecordsTraceRouteHop                     `json:"traceRouteHops,omitempty"`
	WifiBand                   *CallRecordsNetworkInfoWifiBand                 `json:"wifiBand,omitempty"`
	WifiBatteryCharge          *int64                                          `json:"wifiBatteryCharge,omitempty"`
	WifiChannel                *int64                                          `json:"wifiChannel,omitempty"`
	WifiMicrosoftDriver        *string                                         `json:"wifiMicrosoftDriver,omitempty"`
	WifiMicrosoftDriverVersion *string                                         `json:"wifiMicrosoftDriverVersion,omitempty"`
	WifiRadioType              *CallRecordsNetworkInfoWifiRadioType            `json:"wifiRadioType,omitempty"`
	WifiSignalStrength         *int64                                          `json:"wifiSignalStrength,omitempty"`
	WifiVendorDriver           *string                                         `json:"wifiVendorDriver,omitempty"`
	WifiVendorDriverVersion    *string                                         `json:"wifiVendorDriverVersion,omitempty"`
}
