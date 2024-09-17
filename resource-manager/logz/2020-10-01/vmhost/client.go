package vmhost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMHostClient struct {
	Client *resourcemanager.Client
}

func NewVMHostClientWithBaseURI(sdkApi sdkEnv.Api) (*VMHostClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vmhost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMHostClient: %+v", err)
	}

	return &VMHostClient{
		Client: client,
	}, nil
}
