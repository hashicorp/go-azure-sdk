package servicetaskresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTaskResourceClient struct {
	Client *resourcemanager.Client
}

func NewServiceTaskResourceClientWithBaseURI(api sdkEnv.Api) (*ServiceTaskResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "servicetaskresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceTaskResourceClient: %+v", err)
	}

	return &ServiceTaskResourceClient{
		Client: client,
	}, nil
}
