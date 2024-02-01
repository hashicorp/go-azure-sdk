package sqlpoolsrestorepoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsRestorePointsClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsRestorePointsClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsRestorePointsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsrestorepoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsRestorePointsClient: %+v", err)
	}

	return &SqlPoolsRestorePointsClient{
		Client: client,
	}, nil
}
