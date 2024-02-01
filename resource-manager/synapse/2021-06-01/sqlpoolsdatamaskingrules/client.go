package sqlpoolsdatamaskingrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsDataMaskingRulesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsDataMaskingRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsDataMaskingRulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsdatamaskingrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsDataMaskingRulesClient: %+v", err)
	}

	return &SqlPoolsDataMaskingRulesClient{
		Client: client,
	}, nil
}
