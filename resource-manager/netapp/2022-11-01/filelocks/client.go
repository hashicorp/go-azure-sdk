package filelocks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileLocksClient struct {
	Client *resourcemanager.Client
}

func NewFileLocksClientWithBaseURI(api sdkEnv.Api) (*FileLocksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "filelocks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FileLocksClient: %+v", err)
	}

	return &FileLocksClient{
		Client: client,
	}, nil
}
