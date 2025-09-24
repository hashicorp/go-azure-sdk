package v2025_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/afddomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/afdendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/afdorigingroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/afdorigins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/cdnwebapplicationfirewallpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/customdomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/origingroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/origins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/profiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/routes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/rulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/securitypolicies"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AFDDomains                        *afddomains.AFDDomainsClient
	AFDEndpoints                      *afdendpoints.AFDEndpointsClient
	AFDOriginGroups                   *afdorigingroups.AFDOriginGroupsClient
	AFDOrigins                        *afdorigins.AFDOriginsClient
	CdnWebApplicationFirewallPolicies *cdnwebapplicationfirewallpolicies.CdnWebApplicationFirewallPoliciesClient
	CustomDomains                     *customdomains.CustomDomainsClient
	Endpoints                         *endpoints.EndpointsClient
	Openapis                          *openapis.OpenapisClient
	OriginGroups                      *origingroups.OriginGroupsClient
	Origins                           *origins.OriginsClient
	Profiles                          *profiles.ProfilesClient
	Routes                            *routes.RoutesClient
	RuleSets                          *rulesets.RuleSetsClient
	Rules                             *rules.RulesClient
	Secrets                           *secrets.SecretsClient
	SecurityPolicies                  *securitypolicies.SecurityPoliciesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aFDDomainsClient, err := afddomains.NewAFDDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDDomains client: %+v", err)
	}
	configureFunc(aFDDomainsClient.Client)

	aFDEndpointsClient, err := afdendpoints.NewAFDEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDEndpoints client: %+v", err)
	}
	configureFunc(aFDEndpointsClient.Client)

	aFDOriginGroupsClient, err := afdorigingroups.NewAFDOriginGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDOriginGroups client: %+v", err)
	}
	configureFunc(aFDOriginGroupsClient.Client)

	aFDOriginsClient, err := afdorigins.NewAFDOriginsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDOrigins client: %+v", err)
	}
	configureFunc(aFDOriginsClient.Client)

	cdnWebApplicationFirewallPoliciesClient, err := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CdnWebApplicationFirewallPolicies client: %+v", err)
	}
	configureFunc(cdnWebApplicationFirewallPoliciesClient.Client)

	customDomainsClient, err := customdomains.NewCustomDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomDomains client: %+v", err)
	}
	configureFunc(customDomainsClient.Client)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	originGroupsClient, err := origingroups.NewOriginGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OriginGroups client: %+v", err)
	}
	configureFunc(originGroupsClient.Client)

	originsClient, err := origins.NewOriginsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Origins client: %+v", err)
	}
	configureFunc(originsClient.Client)

	profilesClient, err := profiles.NewProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Profiles client: %+v", err)
	}
	configureFunc(profilesClient.Client)

	routesClient, err := routes.NewRoutesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Routes client: %+v", err)
	}
	configureFunc(routesClient.Client)

	ruleSetsClient, err := rulesets.NewRuleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RuleSets client: %+v", err)
	}
	configureFunc(ruleSetsClient.Client)

	rulesClient, err := rules.NewRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Rules client: %+v", err)
	}
	configureFunc(rulesClient.Client)

	secretsClient, err := secrets.NewSecretsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	securityPoliciesClient, err := securitypolicies.NewSecurityPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityPolicies client: %+v", err)
	}
	configureFunc(securityPoliciesClient.Client)

	return &Client{
		AFDDomains:                        aFDDomainsClient,
		AFDEndpoints:                      aFDEndpointsClient,
		AFDOriginGroups:                   aFDOriginGroupsClient,
		AFDOrigins:                        aFDOriginsClient,
		CdnWebApplicationFirewallPolicies: cdnWebApplicationFirewallPoliciesClient,
		CustomDomains:                     customDomainsClient,
		Endpoints:                         endpointsClient,
		Openapis:                          openapisClient,
		OriginGroups:                      originGroupsClient,
		Origins:                           originsClient,
		Profiles:                          profilesClient,
		Routes:                            routesClient,
		RuleSets:                          ruleSetsClient,
		Rules:                             rulesClient,
		Secrets:                           secretsClient,
		SecurityPolicies:                  securityPoliciesClient,
	}, nil
}
