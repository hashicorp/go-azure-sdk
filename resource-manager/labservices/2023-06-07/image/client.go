package image

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageClient struct {
	Client *resourcemanager.Client
}

func NewImageClientWithBaseURI(sdkApi sdkEnv.Api) (*ImageClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "image", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImageClient: %+v", err)
	}

	return &ImageClient{
		Client: client,
	}, nil
}
