package recoverablesqlpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableSqlPoolsClient struct {
	Client *resourcemanager.Client
}

func NewRecoverableSqlPoolsClientWithBaseURI(sdkApi sdkEnv.Api) (*RecoverableSqlPoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "recoverablesqlpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoverableSqlPoolsClient: %+v", err)
	}

	return &RecoverableSqlPoolsClient{
		Client: client,
	}, nil
}
