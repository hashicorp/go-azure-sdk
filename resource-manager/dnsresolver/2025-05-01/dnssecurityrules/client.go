package dnssecurityrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsSecurityRulesClient struct {
	Client *resourcemanager.Client
}

func NewDnsSecurityRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*DnsSecurityRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dnssecurityrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DnsSecurityRulesClient: %+v", err)
	}

	return &DnsSecurityRulesClient{
		Client: client,
	}, nil
}
