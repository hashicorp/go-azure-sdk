package v2025_04_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdcustomdomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdorigingroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdorigins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/checkendpointnameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/checknameavailabilitywithsubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/customdomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/edgenodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/loganalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/origingroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/origins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/profiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/routes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/rulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/securitypolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/validateprobe"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/wafloganalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/webapplicationfirewallmanagedrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/webapplicationfirewallpolicies"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AFDCustomDomains                      *afdcustomdomains.AFDCustomDomainsClient
	AFDEndpoints                          *afdendpoints.AFDEndpointsClient
	AFDOriginGroups                       *afdorigingroups.AFDOriginGroupsClient
	AFDOrigins                            *afdorigins.AFDOriginsClient
	AFDProfiles                           *afdprofiles.AFDProfilesClient
	CheckEndpointNameAvailability         *checkendpointnameavailability.CheckEndpointNameAvailabilityClient
	CheckNameAvailability                 *checknameavailability.CheckNameAvailabilityClient
	CheckNameAvailabilityWithSubscription *checknameavailabilitywithsubscription.CheckNameAvailabilityWithSubscriptionClient
	CustomDomains                         *customdomains.CustomDomainsClient
	Edgenodes                             *edgenodes.EdgenodesClient
	Endpoints                             *endpoints.EndpointsClient
	LogAnalytics                          *loganalytics.LogAnalyticsClient
	OriginGroups                          *origingroups.OriginGroupsClient
	Origins                               *origins.OriginsClient
	Profiles                              *profiles.ProfilesClient
	Routes                                *routes.RoutesClient
	RuleSets                              *rulesets.RuleSetsClient
	Rules                                 *rules.RulesClient
	Secrets                               *secrets.SecretsClient
	SecurityPolicies                      *securitypolicies.SecurityPoliciesClient
	ValidateProbe                         *validateprobe.ValidateProbeClient
	WafLogAnalytics                       *wafloganalytics.WafLogAnalyticsClient
	WebApplicationFirewallManagedRuleSets *webapplicationfirewallmanagedrulesets.WebApplicationFirewallManagedRuleSetsClient
	WebApplicationFirewallPolicies        *webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aFDCustomDomainsClient, err := afdcustomdomains.NewAFDCustomDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDCustomDomains client: %+v", err)
	}
	configureFunc(aFDCustomDomainsClient.Client)

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

	aFDProfilesClient, err := afdprofiles.NewAFDProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AFDProfiles client: %+v", err)
	}
	configureFunc(aFDProfilesClient.Client)

	checkEndpointNameAvailabilityClient, err := checkendpointnameavailability.NewCheckEndpointNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckEndpointNameAvailability client: %+v", err)
	}
	configureFunc(checkEndpointNameAvailabilityClient.Client)

	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	checkNameAvailabilityWithSubscriptionClient, err := checknameavailabilitywithsubscription.NewCheckNameAvailabilityWithSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailabilityWithSubscription client: %+v", err)
	}
	configureFunc(checkNameAvailabilityWithSubscriptionClient.Client)

	customDomainsClient, err := customdomains.NewCustomDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomDomains client: %+v", err)
	}
	configureFunc(customDomainsClient.Client)

	edgenodesClient, err := edgenodes.NewEdgenodesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Edgenodes client: %+v", err)
	}
	configureFunc(edgenodesClient.Client)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	logAnalyticsClient, err := loganalytics.NewLogAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogAnalytics client: %+v", err)
	}
	configureFunc(logAnalyticsClient.Client)

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

	validateProbeClient, err := validateprobe.NewValidateProbeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ValidateProbe client: %+v", err)
	}
	configureFunc(validateProbeClient.Client)

	wafLogAnalyticsClient, err := wafloganalytics.NewWafLogAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WafLogAnalytics client: %+v", err)
	}
	configureFunc(wafLogAnalyticsClient.Client)

	webApplicationFirewallManagedRuleSetsClient, err := webapplicationfirewallmanagedrulesets.NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebApplicationFirewallManagedRuleSets client: %+v", err)
	}
	configureFunc(webApplicationFirewallManagedRuleSetsClient.Client)

	webApplicationFirewallPoliciesClient, err := webapplicationfirewallpolicies.NewWebApplicationFirewallPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebApplicationFirewallPolicies client: %+v", err)
	}
	configureFunc(webApplicationFirewallPoliciesClient.Client)

	return &Client{
		AFDCustomDomains:                      aFDCustomDomainsClient,
		AFDEndpoints:                          aFDEndpointsClient,
		AFDOriginGroups:                       aFDOriginGroupsClient,
		AFDOrigins:                            aFDOriginsClient,
		AFDProfiles:                           aFDProfilesClient,
		CheckEndpointNameAvailability:         checkEndpointNameAvailabilityClient,
		CheckNameAvailability:                 checkNameAvailabilityClient,
		CheckNameAvailabilityWithSubscription: checkNameAvailabilityWithSubscriptionClient,
		CustomDomains:                         customDomainsClient,
		Edgenodes:                             edgenodesClient,
		Endpoints:                             endpointsClient,
		LogAnalytics:                          logAnalyticsClient,
		OriginGroups:                          originGroupsClient,
		Origins:                               originsClient,
		Profiles:                              profilesClient,
		Routes:                                routesClient,
		RuleSets:                              ruleSetsClient,
		Rules:                                 rulesClient,
		Secrets:                               secretsClient,
		SecurityPolicies:                      securityPoliciesClient,
		ValidateProbe:                         validateProbeClient,
		WafLogAnalytics:                       wafLogAnalyticsClient,
		WebApplicationFirewallManagedRuleSets: webApplicationFirewallManagedRuleSetsClient,
		WebApplicationFirewallPolicies:        webApplicationFirewallPoliciesClient,
	}, nil
}
