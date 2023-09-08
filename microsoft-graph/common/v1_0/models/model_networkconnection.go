package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnection struct {
	ApplicationName          *string                     `json:"applicationName,omitempty"`
	DestinationAddress       *string                     `json:"destinationAddress,omitempty"`
	DestinationDomain        *string                     `json:"destinationDomain,omitempty"`
	DestinationLocation      *string                     `json:"destinationLocation,omitempty"`
	DestinationPort          *string                     `json:"destinationPort,omitempty"`
	DestinationUrl           *string                     `json:"destinationUrl,omitempty"`
	Direction                *NetworkConnectionDirection `json:"direction,omitempty"`
	DomainRegisteredDateTime *string                     `json:"domainRegisteredDateTime,omitempty"`
	LocalDnsName             *string                     `json:"localDnsName,omitempty"`
	NatDestinationAddress    *string                     `json:"natDestinationAddress,omitempty"`
	NatDestinationPort       *string                     `json:"natDestinationPort,omitempty"`
	NatSourceAddress         *string                     `json:"natSourceAddress,omitempty"`
	NatSourcePort            *string                     `json:"natSourcePort,omitempty"`
	ODataType                *string                     `json:"@odata.type,omitempty"`
	Protocol                 *NetworkConnectionProtocol  `json:"protocol,omitempty"`
	RiskScore                *string                     `json:"riskScore,omitempty"`
	SourceAddress            *string                     `json:"sourceAddress,omitempty"`
	SourceLocation           *string                     `json:"sourceLocation,omitempty"`
	SourcePort               *string                     `json:"sourcePort,omitempty"`
	Status                   *NetworkConnectionStatus    `json:"status,omitempty"`
	UrlParameters            *string                     `json:"urlParameters,omitempty"`
}
