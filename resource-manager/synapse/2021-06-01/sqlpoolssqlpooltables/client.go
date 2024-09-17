package sqlpoolssqlpooltables

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolTablesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSqlPoolTablesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSqlPoolTablesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolssqlpooltables", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSqlPoolTablesClient: %+v", err)
	}

	return &SqlPoolsSqlPoolTablesClient{
		Client: client,
	}, nil
}
