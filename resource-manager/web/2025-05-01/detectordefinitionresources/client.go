package detectordefinitionresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorDefinitionResourcesClient struct {
	Client *resourcemanager.Client
}

func NewDetectorDefinitionResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*DetectorDefinitionResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "detectordefinitionresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DetectorDefinitionResourcesClient: %+v", err)
	}

	return &DetectorDefinitionResourcesClient{
		Client: client,
	}, nil
}
