package dnsresolverdomainlists

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsResolverDomainListsClient struct {
	Client *resourcemanager.Client
}

func NewDnsResolverDomainListsClientWithBaseURI(sdkApi sdkEnv.Api) (*DnsResolverDomainListsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dnsresolverdomainlists", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DnsResolverDomainListsClient: %+v", err)
	}

	return &DnsResolverDomainListsClient{
		Client: client,
	}, nil
}
