package mongodbtasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoDbTasksClient struct {
	Client *resourcemanager.Client
}

func NewMongoDbTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MongoDbTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "mongodbtasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MongoDbTasksClient: %+v", err)
	}

	return &MongoDbTasksClient{
		Client: client,
	}, nil
}
