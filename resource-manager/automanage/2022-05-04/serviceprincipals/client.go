package serviceprincipals

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalsClient struct {
	Client *resourcemanager.Client
}

func NewServicePrincipalsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServicePrincipalsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "serviceprincipals", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalsClient: %+v", err)
	}

	return &ServicePrincipalsClient{
		Client: client,
	}, nil
}
