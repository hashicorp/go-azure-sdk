package triggeredjobhistories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredJobHistoriesClient struct {
	Client *resourcemanager.Client
}

func NewTriggeredJobHistoriesClientWithBaseURI(sdkApi sdkEnv.Api) (*TriggeredJobHistoriesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "triggeredjobhistories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggeredJobHistoriesClient: %+v", err)
	}

	return &TriggeredJobHistoriesClient{
		Client: client,
	}, nil
}
