// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_07_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/dnsforwardingrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/dnsresolvers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/forwardingrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/inboundendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/outboundendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/virtualnetworklinks"
)

type Client struct {
	DnsForwardingRulesets *dnsforwardingrulesets.DnsForwardingRulesetsClient
	DnsResolvers          *dnsresolvers.DnsResolversClient
	ForwardingRules       *forwardingrules.ForwardingRulesClient
	InboundEndpoints      *inboundendpoints.InboundEndpointsClient
	OutboundEndpoints     *outboundendpoints.OutboundEndpointsClient
	VirtualNetworkLinks   *virtualnetworklinks.VirtualNetworkLinksClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dnsForwardingRulesetsClient := dnsforwardingrulesets.NewDnsForwardingRulesetsClientWithBaseURI(endpoint)
	configureAuthFunc(&dnsForwardingRulesetsClient.Client)

	dnsResolversClient := dnsresolvers.NewDnsResolversClientWithBaseURI(endpoint)
	configureAuthFunc(&dnsResolversClient.Client)

	forwardingRulesClient := forwardingrules.NewForwardingRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&forwardingRulesClient.Client)

	inboundEndpointsClient := inboundendpoints.NewInboundEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&inboundEndpointsClient.Client)

	outboundEndpointsClient := outboundendpoints.NewOutboundEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundEndpointsClient.Client)

	virtualNetworkLinksClient := virtualnetworklinks.NewVirtualNetworkLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkLinksClient.Client)

	return Client{
		DnsForwardingRulesets: &dnsForwardingRulesetsClient,
		DnsResolvers:          &dnsResolversClient,
		ForwardingRules:       &forwardingRulesClient,
		InboundEndpoints:      &inboundEndpointsClient,
		OutboundEndpoints:     &outboundEndpointsClient,
		VirtualNetworkLinks:   &virtualNetworkLinksClient,
	}
}
