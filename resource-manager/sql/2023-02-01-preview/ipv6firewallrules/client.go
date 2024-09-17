package ipv6firewallrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPv6FirewallRulesClient struct {
	Client *resourcemanager.Client
}

func NewIPv6FirewallRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*IPv6FirewallRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "ipv6firewallrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IPv6FirewallRulesClient: %+v", err)
	}

	return &IPv6FirewallRulesClient{
		Client: client,
	}, nil
}
