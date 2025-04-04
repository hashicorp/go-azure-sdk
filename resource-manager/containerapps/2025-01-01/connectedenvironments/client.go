package connectedenvironments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedEnvironmentsClient struct {
	Client *resourcemanager.Client
}

func NewConnectedEnvironmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectedEnvironmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "connectedenvironments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectedEnvironmentsClient: %+v", err)
	}

	return &ConnectedEnvironmentsClient{
		Client: client,
	}, nil
}
