package sqlpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsClient: %+v", err)
	}

	return &SqlPoolsClient{
		Client: client,
	}, nil
}
