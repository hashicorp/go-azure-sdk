package processinfos

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessInfosClient struct {
	Client *resourcemanager.Client
}

func NewProcessInfosClientWithBaseURI(sdkApi sdkEnv.Api) (*ProcessInfosClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "processinfos", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProcessInfosClient: %+v", err)
	}

	return &ProcessInfosClient{
		Client: client,
	}, nil
}
