package restorabledroppedsqlpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedSqlPoolsClient struct {
	Client *resourcemanager.Client
}

func NewRestorableDroppedSqlPoolsClientWithBaseURI(sdkApi sdkEnv.Api) (*RestorableDroppedSqlPoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "restorabledroppedsqlpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestorableDroppedSqlPoolsClient: %+v", err)
	}

	return &RestorableDroppedSqlPoolsClient{
		Client: client,
	}, nil
}
