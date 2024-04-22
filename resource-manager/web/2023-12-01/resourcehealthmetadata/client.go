package resourcehealthmetadata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthMetadataClient struct {
	Client *resourcemanager.Client
}

func NewResourceHealthMetadataClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourceHealthMetadataClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "resourcehealthmetadata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceHealthMetadataClient: %+v", err)
	}

	return &ResourceHealthMetadataClient{
		Client: client,
	}, nil
}
