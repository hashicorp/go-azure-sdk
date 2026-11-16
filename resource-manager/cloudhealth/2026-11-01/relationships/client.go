package relationships

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelationshipsClient struct {
	Client *resourcemanager.Client
}

func NewRelationshipsClientWithBaseURI(sdkApi sdkEnv.Api) (*RelationshipsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "relationships", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RelationshipsClient: %+v", err)
	}

	return &RelationshipsClient{
		Client: client,
	}, nil
}
