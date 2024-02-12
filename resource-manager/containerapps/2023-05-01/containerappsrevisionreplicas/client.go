package containerappsrevisionreplicas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsRevisionReplicasClient struct {
	Client *resourcemanager.Client
}

func NewContainerAppsRevisionReplicasClientWithBaseURI(sdkApi sdkEnv.Api) (*ContainerAppsRevisionReplicasClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "containerappsrevisionreplicas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainerAppsRevisionReplicasClient: %+v", err)
	}

	return &ContainerAppsRevisionReplicasClient{
		Client: client,
	}, nil
}
