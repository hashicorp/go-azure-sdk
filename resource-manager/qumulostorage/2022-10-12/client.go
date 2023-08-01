package v2022_10_12

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/qumulostorage/2022-10-12/filesystems"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	FileSystems *filesystems.FileSystemsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	fileSystemsClient, err := filesystems.NewFileSystemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileSystems client: %+v", err)
	}
	configureFunc(fileSystemsClient.Client)

	return &Client{
		FileSystems: fileSystemsClient,
	}, nil
}
