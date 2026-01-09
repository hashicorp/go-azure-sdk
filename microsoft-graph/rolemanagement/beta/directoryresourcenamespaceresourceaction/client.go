package directoryresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewDirectoryResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryResourceNamespaceResourceActionClient: %+v", err)
	}

	return &DirectoryResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
