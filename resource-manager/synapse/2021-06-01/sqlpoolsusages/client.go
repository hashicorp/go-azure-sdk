package sqlpoolsusages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsUsagesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsUsagesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsUsagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolsusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsUsagesClient: %+v", err)
	}

	return &SqlPoolsUsagesClient{
		Client: client,
	}, nil
}
