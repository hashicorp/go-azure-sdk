package asyncoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsyncOperationsClient struct {
	Client *resourcemanager.Client
}

func NewAsyncOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*AsyncOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "asyncoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AsyncOperationsClient: %+v", err)
	}

	return &AsyncOperationsClient{
		Client: client,
	}, nil
}
