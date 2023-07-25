package customimages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImagesClient struct {
	Client *resourcemanager.Client
}

func NewCustomImagesClientWithBaseURI(api environments.Api) (*CustomImagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "customimages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomImagesClient: %+v", err)
	}

	return &CustomImagesClient{
		Client: client,
	}, nil
}
