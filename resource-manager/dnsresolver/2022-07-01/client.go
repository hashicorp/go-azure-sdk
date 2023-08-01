package v2022_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/dnsforwardingrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/dnsresolvers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/forwardingrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/inboundendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/outboundendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2022-07-01/virtualnetworklinks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DnsForwardingRulesets *dnsforwardingrulesets.DnsForwardingRulesetsClient
	DnsResolvers          *dnsresolvers.DnsResolversClient
	ForwardingRules       *forwardingrules.ForwardingRulesClient
	InboundEndpoints      *inboundendpoints.InboundEndpointsClient
	OutboundEndpoints     *outboundendpoints.OutboundEndpointsClient
	VirtualNetworkLinks   *virtualnetworklinks.VirtualNetworkLinksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dnsForwardingRulesetsClient, err := dnsforwardingrulesets.NewDnsForwardingRulesetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DnsForwardingRulesets client: %+v", err)
	}
	configureFunc(dnsForwardingRulesetsClient.Client)

	dnsResolversClient, err := dnsresolvers.NewDnsResolversClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DnsResolvers client: %+v", err)
	}
	configureFunc(dnsResolversClient.Client)

	forwardingRulesClient, err := forwardingrules.NewForwardingRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ForwardingRules client: %+v", err)
	}
	configureFunc(forwardingRulesClient.Client)

	inboundEndpointsClient, err := inboundendpoints.NewInboundEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InboundEndpoints client: %+v", err)
	}
	configureFunc(inboundEndpointsClient.Client)

	outboundEndpointsClient, err := outboundendpoints.NewOutboundEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundEndpoints client: %+v", err)
	}
	configureFunc(outboundEndpointsClient.Client)

	virtualNetworkLinksClient, err := virtualnetworklinks.NewVirtualNetworkLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkLinks client: %+v", err)
	}
	configureFunc(virtualNetworkLinksClient.Client)

	return &Client{
		DnsForwardingRulesets: dnsForwardingRulesetsClient,
		DnsResolvers:          dnsResolversClient,
		ForwardingRules:       forwardingRulesClient,
		InboundEndpoints:      inboundEndpointsClient,
		OutboundEndpoints:     outboundEndpointsClient,
		VirtualNetworkLinks:   virtualNetworkLinksClient,
	}, nil
}
