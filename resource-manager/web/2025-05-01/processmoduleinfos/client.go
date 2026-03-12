package processmoduleinfos

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessModuleInfosClient struct {
	Client *resourcemanager.Client
}

func NewProcessModuleInfosClientWithBaseURI(sdkApi sdkEnv.Api) (*ProcessModuleInfosClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "processmoduleinfos", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProcessModuleInfosClient: %+v", err)
	}

	return &ProcessModuleInfosClient{
		Client: client,
	}, nil
}
