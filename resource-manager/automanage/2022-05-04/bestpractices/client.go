package bestpractices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BestPracticesClient struct {
	Client *resourcemanager.Client
}

func NewBestPracticesClientWithBaseURI(sdkApi sdkEnv.Api) (*BestPracticesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bestpractices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BestPracticesClient: %+v", err)
	}

	return &BestPracticesClient{
		Client: client,
	}, nil
}
