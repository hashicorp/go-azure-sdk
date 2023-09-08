package directory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryClient struct {
	Client *msgraph.Client
}

func NewDirectoryClientWithBaseURI(api sdkEnv.Api) (*DirectoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryClient: %+v", err)
	}

	return &DirectoryClient{
		Client: client,
	}, nil
}
