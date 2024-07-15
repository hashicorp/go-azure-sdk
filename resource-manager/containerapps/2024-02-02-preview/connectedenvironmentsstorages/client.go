package connectedenvironmentsstorages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedEnvironmentsStoragesClient struct {
	Client *resourcemanager.Client
}

func NewConnectedEnvironmentsStoragesClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectedEnvironmentsStoragesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connectedenvironmentsstorages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectedEnvironmentsStoragesClient: %+v", err)
	}

	return &ConnectedEnvironmentsStoragesClient{
		Client: client,
	}, nil
}
