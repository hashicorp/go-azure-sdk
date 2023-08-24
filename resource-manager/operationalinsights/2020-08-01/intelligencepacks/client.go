package intelligencepacks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntelligencePacksClient struct {
	Client *resourcemanager.Client
}

func NewIntelligencePacksClientWithBaseURI(sdkApi sdkEnv.Api) (*IntelligencePacksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "intelligencepacks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntelligencePacksClient: %+v", err)
	}

	return &IntelligencePacksClient{
		Client: client,
	}, nil
}
