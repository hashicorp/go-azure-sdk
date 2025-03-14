package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopicSpacesConfiguration struct {
	CustomDomains                              *[]CustomDomainConfiguration   `json:"customDomains,omitempty"`
	Hostname                                   *string                        `json:"hostname,omitempty"`
	MaximumClientSessionsPerAuthenticationName *int64                         `json:"maximumClientSessionsPerAuthenticationName,omitempty"`
	MaximumSessionExpiryInHours                *int64                         `json:"maximumSessionExpiryInHours,omitempty"`
	RouteTopicResourceId                       *string                        `json:"routeTopicResourceId,omitempty"`
	RoutingEnrichments                         *RoutingEnrichments            `json:"routingEnrichments,omitempty"`
	RoutingIdentityInfo                        *RoutingIdentityInfo           `json:"routingIdentityInfo,omitempty"`
	State                                      *TopicSpacesConfigurationState `json:"state,omitempty"`
}
