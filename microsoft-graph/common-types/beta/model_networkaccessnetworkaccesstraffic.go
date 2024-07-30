package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTraffic struct {
	Action                       *NetworkaccessNetworkAccessTrafficAction            `json:"action,omitempty"`
	AgentVersion                 *string                                             `json:"agentVersion,omitempty"`
	ConnectionId                 *string                                             `json:"connectionId,omitempty"`
	CreatedDateTime              *string                                             `json:"createdDateTime,omitempty"`
	DestinationFQDN              *string                                             `json:"destinationFQDN,omitempty"`
	DestinationIp                *string                                             `json:"destinationIp,omitempty"`
	DestinationPort              *int64                                              `json:"destinationPort,omitempty"`
	DestinationWebCategory       *NetworkaccessWebCategory                           `json:"destinationWebCategory,omitempty"`
	DeviceCategory               *NetworkaccessNetworkAccessTrafficDeviceCategory    `json:"deviceCategory,omitempty"`
	DeviceId                     *string                                             `json:"deviceId,omitempty"`
	DeviceOperatingSystem        *string                                             `json:"deviceOperatingSystem,omitempty"`
	DeviceOperatingSystemVersion *string                                             `json:"deviceOperatingSystemVersion,omitempty"`
	FilteringProfileId           *string                                             `json:"filteringProfileId,omitempty"`
	FilteringProfileName         *string                                             `json:"filteringProfileName,omitempty"`
	Headers                      *NetworkaccessHeaders                               `json:"headers,omitempty"`
	InitiatingProcessName        *string                                             `json:"initiatingProcessName,omitempty"`
	NetworkProtocol              *NetworkaccessNetworkAccessTrafficNetworkProtocol   `json:"networkProtocol,omitempty"`
	ODataType                    *string                                             `json:"@odata.type,omitempty"`
	PolicyId                     *string                                             `json:"policyId,omitempty"`
	PolicyName                   *string                                             `json:"policyName,omitempty"`
	PolicyRuleId                 *string                                             `json:"policyRuleId,omitempty"`
	PolicyRuleName               *string                                             `json:"policyRuleName,omitempty"`
	ReceivedBytes                *int64                                              `json:"receivedBytes,omitempty"`
	ResourceTenantId             *string                                             `json:"resourceTenantId,omitempty"`
	SentBytes                    *int64                                              `json:"sentBytes,omitempty"`
	SessionId                    *string                                             `json:"sessionId,omitempty"`
	SourceIp                     *string                                             `json:"sourceIp,omitempty"`
	SourcePort                   *int64                                              `json:"sourcePort,omitempty"`
	TenantId                     *string                                             `json:"tenantId,omitempty"`
	TrafficType                  *NetworkaccessNetworkAccessTrafficTrafficType       `json:"trafficType,omitempty"`
	TransactionId                *string                                             `json:"transactionId,omitempty"`
	TransportProtocol            *NetworkaccessNetworkAccessTrafficTransportProtocol `json:"transportProtocol,omitempty"`
	UserId                       *string                                             `json:"userId,omitempty"`
	UserPrincipalName            *string                                             `json:"userPrincipalName,omitempty"`
}
