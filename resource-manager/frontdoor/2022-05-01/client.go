// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2022-05-01/webapplicationfirewallmanagedrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2022-05-01/webapplicationfirewallpolicies"
)

type Client struct {
	WebApplicationFirewallManagedRuleSets *webapplicationfirewallmanagedrulesets.WebApplicationFirewallManagedRuleSetsClient
	WebApplicationFirewallPolicies        *webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	webApplicationFirewallManagedRuleSetsClient := webapplicationfirewallmanagedrulesets.NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallManagedRuleSetsClient.Client)

	webApplicationFirewallPoliciesClient := webapplicationfirewallpolicies.NewWebApplicationFirewallPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallPoliciesClient.Client)

	return Client{
		WebApplicationFirewallManagedRuleSets: &webApplicationFirewallManagedRuleSetsClient,
		WebApplicationFirewallPolicies:        &webApplicationFirewallPoliciesClient,
	}
}
