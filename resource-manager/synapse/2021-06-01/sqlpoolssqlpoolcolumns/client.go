package sqlpoolssqlpoolcolumns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolColumnsClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSqlPoolColumnsClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSqlPoolColumnsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolssqlpoolcolumns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSqlPoolColumnsClient: %+v", err)
	}

	return &SqlPoolsSqlPoolColumnsClient{
		Client: client,
	}, nil
}
