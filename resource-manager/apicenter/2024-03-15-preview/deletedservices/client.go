package deletedservices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServicesClient struct {
	Client *resourcemanager.Client
}

func NewDeletedServicesClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedServicesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "deletedservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedServicesClient: %+v", err)
	}

	return &DeletedServicesClient{
		Client: client,
	}, nil
}
