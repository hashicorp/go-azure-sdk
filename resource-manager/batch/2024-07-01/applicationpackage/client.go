package applicationpackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackageClient struct {
	Client *resourcemanager.Client
}

func NewApplicationPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationPackageClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "applicationpackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationPackageClient: %+v", err)
	}

	return &ApplicationPackageClient{
		Client: client,
	}, nil
}
