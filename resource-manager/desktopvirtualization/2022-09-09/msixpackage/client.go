package msixpackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSIXPackageClient struct {
	Client *resourcemanager.Client
}

func NewMSIXPackageClientWithBaseURI(api environments.Api) (*MSIXPackageClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "msixpackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MSIXPackageClient: %+v", err)
	}

	return &MSIXPackageClient{
		Client: client,
	}, nil
}
