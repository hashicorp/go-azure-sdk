package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/beta/directoryobject"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DirectoryObject *directoryobject.DirectoryObjectClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryObject client: %+v", err)
	}
	configureFunc(directoryObjectClient.Client)

	return &Client{
		DirectoryObject: directoryObjectClient,
	}, nil
}
