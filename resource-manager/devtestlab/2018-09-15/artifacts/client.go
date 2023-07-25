package artifacts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactsClient struct {
	Client *resourcemanager.Client
}

func NewArtifactsClientWithBaseURI(api environments.Api) (*ArtifactsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "artifacts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ArtifactsClient: %+v", err)
	}

	return &ArtifactsClient{
		Client: client,
	}, nil
}
