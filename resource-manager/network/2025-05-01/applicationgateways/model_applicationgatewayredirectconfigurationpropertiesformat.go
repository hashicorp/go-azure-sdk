package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayRedirectConfigurationPropertiesFormat struct {
	IncludePath         *bool                           `json:"includePath,omitempty"`
	IncludeQueryString  *bool                           `json:"includeQueryString,omitempty"`
	PathRules           *[]CommonSubResource            `json:"pathRules,omitempty"`
	RedirectType        *ApplicationGatewayRedirectType `json:"redirectType,omitempty"`
	RequestRoutingRules *[]CommonSubResource            `json:"requestRoutingRules,omitempty"`
	TargetListener      *CommonSubResource              `json:"targetListener,omitempty"`
	TargetURL           *string                         `json:"targetUrl,omitempty"`
	UrlPathMaps         *[]CommonSubResource            `json:"urlPathMaps,omitempty"`
}
