package mobilenetwork

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileNetworkClient struct {
	Client *resourcemanager.Client
}

func NewMobileNetworkClientWithBaseURI(api sdkEnv.Api) (*MobileNetworkClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "mobilenetwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileNetworkClient: %+v", err)
	}

	return &MobileNetworkClient{
		Client: client,
	}, nil
}
