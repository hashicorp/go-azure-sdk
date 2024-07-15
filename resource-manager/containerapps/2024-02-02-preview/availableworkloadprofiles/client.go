package availableworkloadprofiles

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableWorkloadProfilesClient struct {
	Client *resourcemanager.Client
}

func NewAvailableWorkloadProfilesClientWithBaseURI(sdkApi sdkEnv.Api) (*AvailableWorkloadProfilesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "availableworkloadprofiles", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AvailableWorkloadProfilesClient: %+v", err)
	}

	return &AvailableWorkloadProfilesClient{
		Client: client,
	}, nil
}
