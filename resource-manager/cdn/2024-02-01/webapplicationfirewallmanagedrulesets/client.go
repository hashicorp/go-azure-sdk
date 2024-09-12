package webapplicationfirewallmanagedrulesets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplicationFirewallManagedRuleSetsClient struct {
	Client *resourcemanager.Client
}

func NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI(sdkApi sdkEnv.Api) (*WebApplicationFirewallManagedRuleSetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "webapplicationfirewallmanagedrulesets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebApplicationFirewallManagedRuleSetsClient: %+v", err)
	}

	return &WebApplicationFirewallManagedRuleSetsClient{
		Client: client,
	}, nil
}
