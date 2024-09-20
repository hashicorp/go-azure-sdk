package allowedconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewAllowedConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*AllowedConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "allowedconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AllowedConnectionsClient: %+v", err)
	}

	return &AllowedConnectionsClient{
		Client: client,
	}, nil
}
