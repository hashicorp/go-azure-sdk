package incrementalrestorepoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncrementalRestorePointsClient struct {
	Client *resourcemanager.Client
}

func NewIncrementalRestorePointsClientWithBaseURI(sdkApi sdkEnv.Api) (*IncrementalRestorePointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "incrementalrestorepoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncrementalRestorePointsClient: %+v", err)
	}

	return &IncrementalRestorePointsClient{
		Client: client,
	}, nil
}
