package vmcollectionupdate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMCollectionUpdateClient struct {
	Client *resourcemanager.Client
}

func NewVMCollectionUpdateClientWithBaseURI(api sdkEnv.Api) (*VMCollectionUpdateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vmcollectionupdate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMCollectionUpdateClient: %+v", err)
	}

	return &VMCollectionUpdateClient{
		Client: client,
	}, nil
}
