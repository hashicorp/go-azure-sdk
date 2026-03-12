package requesthistories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestHistoriesClient struct {
	Client *resourcemanager.Client
}

func NewRequestHistoriesClientWithBaseURI(sdkApi sdkEnv.Api) (*RequestHistoriesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "requesthistories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RequestHistoriesClient: %+v", err)
	}

	return &RequestHistoriesClient{
		Client: client,
	}, nil
}
