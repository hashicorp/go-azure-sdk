package sqlpoolsconnectionpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolConnectionPolicyProperties struct {
	ProxyDnsName          *string `json:"proxyDnsName,omitempty"`
	ProxyPort             *string `json:"proxyPort,omitempty"`
	RedirectionState      *string `json:"redirectionState,omitempty"`
	SecurityEnabledAccess *string `json:"securityEnabledAccess,omitempty"`
	State                 *string `json:"state,omitempty"`
	UseServerDefault      *string `json:"useServerDefault,omitempty"`
	Visibility            *string `json:"visibility,omitempty"`
}
