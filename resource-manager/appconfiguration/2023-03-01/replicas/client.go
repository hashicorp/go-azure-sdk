package replicas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicasClient struct {
	Client *resourcemanager.Client
}

func NewReplicasClientWithBaseURI(api sdkEnv.Api) (*ReplicasClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "replicas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicasClient: %+v", err)
	}

	return &ReplicasClient{
		Client: client,
	}, nil
}
