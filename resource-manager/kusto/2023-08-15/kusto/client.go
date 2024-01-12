package kusto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KustoClient struct {
	Client *resourcemanager.Client
}

func NewKustoClientWithBaseURI(sdkApi sdkEnv.Api) (*KustoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "kusto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KustoClient: %+v", err)
	}

	return &KustoClient{
		Client: client,
	}, nil
}
