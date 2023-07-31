package shares

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharesClient struct {
	Client *resourcemanager.Client
}

func NewSharesClientWithBaseURI(api sdkEnv.Api) (*SharesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "shares", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SharesClient: %+v", err)
	}

	return &SharesClient{
		Client: client,
	}, nil
}
