package newrelics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NewRelicsClient struct {
	Client *resourcemanager.Client
}

func NewNewRelicsClientWithBaseURI(sdkApi sdkEnv.Api) (*NewRelicsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "newrelics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NewRelicsClient: %+v", err)
	}

	return &NewRelicsClient{
		Client: client,
	}, nil
}
