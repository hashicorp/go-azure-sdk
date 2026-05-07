package cdnwebapplicationfirewallpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CdnWebApplicationFirewallPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewCdnWebApplicationFirewallPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*CdnWebApplicationFirewallPoliciesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cdnwebapplicationfirewallpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CdnWebApplicationFirewallPoliciesClient: %+v", err)
	}

	return &CdnWebApplicationFirewallPoliciesClient{
		Client: client,
	}, nil
}
