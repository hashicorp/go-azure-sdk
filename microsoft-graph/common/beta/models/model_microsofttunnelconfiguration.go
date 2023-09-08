package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelConfiguration struct {
	AdvancedSettings      *[]KeyValuePair `json:"advancedSettings,omitempty"`
	DefaultDomainSuffix   *string         `json:"defaultDomainSuffix,omitempty"`
	Description           *string         `json:"description,omitempty"`
	DisableUdpConnections *bool           `json:"disableUdpConnections,omitempty"`
	DisplayName           *string         `json:"displayName,omitempty"`
	DnsServers            *[]string       `json:"dnsServers,omitempty"`
	Id                    *string         `json:"id,omitempty"`
	LastUpdateDateTime    *string         `json:"lastUpdateDateTime,omitempty"`
	ListenPort            *int64          `json:"listenPort,omitempty"`
	Network               *string         `json:"network,omitempty"`
	ODataType             *string         `json:"@odata.type,omitempty"`
	RoleScopeTagIds       *[]string       `json:"roleScopeTagIds,omitempty"`
	RouteExcludes         *[]string       `json:"routeExcludes,omitempty"`
	RouteIncludes         *[]string       `json:"routeIncludes,omitempty"`
	RoutesExclude         *[]string       `json:"routesExclude,omitempty"`
	RoutesInclude         *[]string       `json:"routesInclude,omitempty"`
	SplitDNS              *[]string       `json:"splitDNS,omitempty"`
}
