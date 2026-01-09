package directoryresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewDirectoryResourceNamespaceClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryResourceNamespaceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryResourceNamespaceClient: %+v", err)
	}

	return &DirectoryResourceNamespaceClient{
		Client: client,
	}, nil
}
