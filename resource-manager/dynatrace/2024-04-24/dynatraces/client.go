package dynatraces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DynatracesClient struct {
	Client *resourcemanager.Client
}

func NewDynatracesClientWithBaseURI(sdkApi sdkEnv.Api) (*DynatracesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dynatraces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DynatracesClient: %+v", err)
	}

	return &DynatracesClient{
		Client: client,
	}, nil
}
