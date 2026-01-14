package applicationpackages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackagesClient struct {
	Client *resourcemanager.Client
}

func NewApplicationPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationPackagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "applicationpackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationPackagesClient: %+v", err)
	}

	return &ApplicationPackagesClient{
		Client: client,
	}, nil
}
