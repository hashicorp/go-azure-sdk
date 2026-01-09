package directoryresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &DirectoryResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
