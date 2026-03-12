package premieraddonoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PremierAddOnOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewPremierAddOnOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PremierAddOnOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "premieraddonoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PremierAddOnOperationGroupClient: %+v", err)
	}

	return &PremierAddOnOperationGroupClient{
		Client: client,
	}, nil
}
