package serveradvisors

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerAdvisorsClient struct {
	Client *resourcemanager.Client
}

func NewServerAdvisorsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerAdvisorsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "serveradvisors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerAdvisorsClient: %+v", err)
	}

	return &ServerAdvisorsClient{
		Client: client,
	}, nil
}
