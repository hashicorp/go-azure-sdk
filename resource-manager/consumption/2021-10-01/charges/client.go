package charges

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargesClient struct {
	Client *resourcemanager.Client
}

func NewChargesClientWithBaseURI(api sdkEnv.Api) (*ChargesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "charges", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChargesClient: %+v", err)
	}

	return &ChargesClient{
		Client: client,
	}, nil
}
