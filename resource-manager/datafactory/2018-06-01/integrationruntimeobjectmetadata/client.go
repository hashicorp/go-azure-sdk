package integrationruntimeobjectmetadata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeObjectMetadataClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationRuntimeObjectMetadataClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationRuntimeObjectMetadataClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "integrationruntimeobjectmetadata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationRuntimeObjectMetadataClient: %+v", err)
	}

	return &IntegrationRuntimeObjectMetadataClient{
		Client: client,
	}, nil
}
