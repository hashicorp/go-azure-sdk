package resourcehealthmetadataoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthMetadataOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewResourceHealthMetadataOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourceHealthMetadataOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "resourcehealthmetadataoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceHealthMetadataOperationGroupClient: %+v", err)
	}

	return &ResourceHealthMetadataOperationGroupClient{
		Client: client,
	}, nil
}
