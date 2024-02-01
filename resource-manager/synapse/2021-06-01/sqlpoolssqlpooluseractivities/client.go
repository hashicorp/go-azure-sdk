package sqlpoolssqlpooluseractivities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolUserActivitiesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSqlPoolUserActivitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSqlPoolUserActivitiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolssqlpooluseractivities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSqlPoolUserActivitiesClient: %+v", err)
	}

	return &SqlPoolsSqlPoolUserActivitiesClient{
		Client: client,
	}, nil
}
