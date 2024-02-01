package sqlpoolsoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsOperationsClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsOperationsClient: %+v", err)
	}

	return &SqlPoolsOperationsClient{
		Client: client,
	}, nil
}
