package appserviceenvironments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsClient struct {
	Client *resourcemanager.Client
}

func NewAppServiceEnvironmentsClientWithBaseURI(api environments.Api) (*AppServiceEnvironmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "appserviceenvironments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppServiceEnvironmentsClient: %+v", err)
	}

	return &AppServiceEnvironmentsClient{
		Client: client,
	}, nil
}
