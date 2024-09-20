package applicationtype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTypeClient struct {
	Client *resourcemanager.Client
}

func NewApplicationTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationTypeClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "applicationtype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationTypeClient: %+v", err)
	}

	return &ApplicationTypeClient{
		Client: client,
	}, nil
}
