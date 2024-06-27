package transfers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransfersClient struct {
	Client *resourcemanager.Client
}

func NewTransfersClientWithBaseURI(sdkApi sdkEnv.Api) (*TransfersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "transfers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransfersClient: %+v", err)
	}

	return &TransfersClient{
		Client: client,
	}, nil
}
