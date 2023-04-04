package hybrididentitymetadata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridIdentityMetadataClient struct {
	Client *resourcemanager.Client
}

func NewHybridIdentityMetadataClientWithBaseURI(api environments.Api) (*HybridIdentityMetadataClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hybrididentitymetadata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridIdentityMetadataClient: %+v", err)
	}

	return &HybridIdentityMetadataClient{
		Client: client,
	}, nil
}
