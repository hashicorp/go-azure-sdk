package dppjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppJobClient struct {
	Client *resourcemanager.Client
}

func NewDppJobClientWithBaseURI(sdkApi sdkEnv.Api) (*DppJobClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dppjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DppJobClient: %+v", err)
	}

	return &DppJobClient{
		Client: client,
	}, nil
}
