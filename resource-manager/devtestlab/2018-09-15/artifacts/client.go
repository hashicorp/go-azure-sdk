package artifacts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactsClient struct {
	Client *resourcemanager.Client
}

func NewArtifactsClientWithBaseURI(sdkApi sdkEnv.Api) (*ArtifactsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "artifacts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ArtifactsClient: %+v", err)
	}

	return &ArtifactsClient{
		Client: client,
	}, nil
}
