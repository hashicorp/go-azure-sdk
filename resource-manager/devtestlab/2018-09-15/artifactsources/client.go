package artifactsources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactSourcesClient struct {
	Client *resourcemanager.Client
}

func NewArtifactSourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*ArtifactSourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "artifactsources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ArtifactSourcesClient: %+v", err)
	}

	return &ArtifactSourcesClient{
		Client: client,
	}, nil
}
