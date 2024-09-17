package datatransfer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataTransferClient struct {
	Client *resourcemanager.Client
}

func NewDataTransferClientWithBaseURI(sdkApi sdkEnv.Api) (*DataTransferClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "datatransfer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataTransferClient: %+v", err)
	}

	return &DataTransferClient{
		Client: client,
	}, nil
}
