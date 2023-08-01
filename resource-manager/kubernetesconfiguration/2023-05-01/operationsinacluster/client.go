package operationsinacluster

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsInAClusterClient struct {
	Client *resourcemanager.Client
}

func NewOperationsInAClusterClientWithBaseURI(sdkApi sdkEnv.Api) (*OperationsInAClusterClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "operationsinacluster", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationsInAClusterClient: %+v", err)
	}

	return &OperationsInAClusterClient{
		Client: client,
	}, nil
}
