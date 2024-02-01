package sqlpoolstables

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsTablesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsTablesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsTablesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolstables", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsTablesClient: %+v", err)
	}

	return &SqlPoolsTablesClient{
		Client: client,
	}, nil
}
