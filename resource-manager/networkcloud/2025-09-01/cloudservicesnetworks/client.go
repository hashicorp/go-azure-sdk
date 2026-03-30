package cloudservicesnetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudServicesNetworksClient struct {
	Client *resourcemanager.Client
}

func NewCloudServicesNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudServicesNetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudservicesnetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudServicesNetworksClient: %+v", err)
	}

	return &CloudServicesNetworksClient{
		Client: client,
	}, nil
}
