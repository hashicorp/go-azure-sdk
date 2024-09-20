package fetchsecondaryrecoverypoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FetchSecondaryRecoveryPointsClient struct {
	Client *resourcemanager.Client
}

func NewFetchSecondaryRecoveryPointsClientWithBaseURI(sdkApi sdkEnv.Api) (*FetchSecondaryRecoveryPointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "fetchsecondaryrecoverypoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FetchSecondaryRecoveryPointsClient: %+v", err)
	}

	return &FetchSecondaryRecoveryPointsClient{
		Client: client,
	}, nil
}
