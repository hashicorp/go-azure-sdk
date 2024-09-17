package databaseautomatictuning

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseAutomaticTuningClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseAutomaticTuningClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseAutomaticTuningClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databaseautomatictuning", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseAutomaticTuningClient: %+v", err)
	}

	return &DatabaseAutomaticTuningClient{
		Client: client,
	}, nil
}
