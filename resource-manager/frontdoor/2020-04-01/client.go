// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/checkfrontdoornameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/checkfrontdoornameavailabilitywithsubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/frontdoors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/webapplicationfirewallmanagedrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/webapplicationfirewallpolicies"
)

type Client struct {
	CheckFrontDoorNameAvailability                 *checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient
	CheckFrontDoorNameAvailabilityWithSubscription *checkfrontdoornameavailabilitywithsubscription.CheckFrontDoorNameAvailabilityWithSubscriptionClient
	FrontDoors                                     *frontdoors.FrontDoorsClient
	WebApplicationFirewallManagedRuleSets          *webapplicationfirewallmanagedrulesets.WebApplicationFirewallManagedRuleSetsClient
	WebApplicationFirewallPolicies                 *webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkFrontDoorNameAvailabilityClient := checkfrontdoornameavailability.NewCheckFrontDoorNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityClient.Client)

	checkFrontDoorNameAvailabilityWithSubscriptionClient := checkfrontdoornameavailabilitywithsubscription.NewCheckFrontDoorNameAvailabilityWithSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityWithSubscriptionClient.Client)

	frontDoorsClient := frontdoors.NewFrontDoorsClientWithBaseURI(endpoint)
	configureAuthFunc(&frontDoorsClient.Client)

	webApplicationFirewallManagedRuleSetsClient := webapplicationfirewallmanagedrulesets.NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallManagedRuleSetsClient.Client)

	webApplicationFirewallPoliciesClient := webapplicationfirewallpolicies.NewWebApplicationFirewallPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallPoliciesClient.Client)

	return Client{
		CheckFrontDoorNameAvailability:                 &checkFrontDoorNameAvailabilityClient,
		CheckFrontDoorNameAvailabilityWithSubscription: &checkFrontDoorNameAvailabilityWithSubscriptionClient,
		FrontDoors:                            &frontDoorsClient,
		WebApplicationFirewallManagedRuleSets: &webApplicationFirewallManagedRuleSetsClient,
		WebApplicationFirewallPolicies:        &webApplicationFirewallPoliciesClient,
	}
}
