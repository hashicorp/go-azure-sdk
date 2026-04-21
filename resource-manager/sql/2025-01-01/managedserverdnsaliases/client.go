package managedserverdnsaliases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedServerDnsAliasesClient struct {
	Client *resourcemanager.Client
}

func NewManagedServerDnsAliasesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedServerDnsAliasesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedserverdnsaliases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedServerDnsAliasesClient: %+v", err)
	}

	return &ManagedServerDnsAliasesClient{
		Client: client,
	}, nil
}
