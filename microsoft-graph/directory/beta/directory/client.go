package directory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryClient struct {
	Client *msgraph.Client
}

func NewDirectoryClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryClient: %+v", err)
	}

	return &DirectoryClient{
		Client: client,
	}, nil
}
