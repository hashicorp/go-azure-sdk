package v2022_12_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2022-12-01-preview/defenderforstorage"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DefenderForStorage *defenderforstorage.DefenderForStorageClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	defenderForStorageClient, err := defenderforstorage.NewDefenderForStorageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderForStorage client: %+v", err)
	}
	configureFunc(defenderForStorageClient.Client)

	return &Client{
		DefenderForStorage: defenderForStorageClient,
	}, nil
}
