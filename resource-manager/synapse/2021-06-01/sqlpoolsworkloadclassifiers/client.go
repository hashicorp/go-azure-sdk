package sqlpoolsworkloadclassifiers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsWorkloadClassifiersClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsWorkloadClassifiersClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsWorkloadClassifiersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsworkloadclassifiers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsWorkloadClassifiersClient: %+v", err)
	}

	return &SqlPoolsWorkloadClassifiersClient{
		Client: client,
	}, nil
}
